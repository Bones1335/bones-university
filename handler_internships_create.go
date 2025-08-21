package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handlerCreateInternships(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		BusinessName        string `json:"business_name"`
		NumSpots            int16  `json:"num_spots"`
		BusinessAddress     string `json:"business_address"`
		BusinessCity        string `json:"business_city"`
		BusinessPostalCode  int    `json:"business_postal_code"`
		BusinessState       string `json:"business_state"`
		BusinessCountry     string `json:"business_country"`
		BusinessPhoneNumber string `json:"business_phone_number"`
		BusinessEmail       string `json:"business_email"`
		BusinessType        string `json:"business_type"`
	}

	type response struct {
		database.Internship
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't validate JWT", err)
		return
	}

	usersRole, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}
	if usersRole != "administrator" && usersRole != "professor" {
		respondWithError(w, http.StatusUnauthorized, "user not authorized to perform this actions", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode internship parameters", err)
		return
	}

	internship, err := cfg.db.CreateInternships(r.Context(), database.CreateInternshipsParams{
		BusinessName:        params.BusinessName,
		NumSpots:            int16(params.NumSpots),
		BusinessAddress:     params.BusinessAddress,
		BusinessCity:        params.BusinessCity,
		BusinessPostalCode:  int32(params.BusinessPostalCode),
		BusinessState:       params.BusinessState,
		BusinessCountry:     params.BusinessCountry,
		BusinessPhoneNumber: params.BusinessPhoneNumber,
		BusinessEmail:       params.BusinessEmail,
		BusinessType:        params.BusinessType,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create internship", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{internship})
}
