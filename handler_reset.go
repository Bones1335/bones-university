package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Forbidden"))
	}

	cfg.db.ResetUsers(r.Context())
	cfg.db.ResetRoles(r.Context())
	cfg.db.ResetUsersRoles(r.Context())
	cfg.db.ResetDegrees(r.Context())
	cfg.db.ResetYears(r.Context())
	cfg.db.ResetUsersPrograms(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("database reset to initial state\n"))
}
