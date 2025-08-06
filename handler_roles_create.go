package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handlerCreateRoles(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		RoleName string `json:"role_name"`
	}

	type response struct {
		database.Role
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode role parameters", err)
		return
	}

	role, err := cfg.db.CreateRole(r.Context(), params.RoleName)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create role", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		role,
	})
}
