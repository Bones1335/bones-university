package main

import "net/http"

func (cfg *apiConfig) handlerGetDegrees(w http.ResponseWriter, r *http.Request) {
	dbDegrees, err := cfg.db.GetDegrees(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get list of degree programs from the database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, dbDegrees)
}
