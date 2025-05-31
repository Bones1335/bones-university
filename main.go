package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bones1335/bones-university/internal/database"
	"github.com/Bones1335/bones-university/internal/env"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	env.SetEnv(".env")

	dbURL := os.Getenv("DB_URL")

	const filepathRoot = "."
	const port = "8080"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Error connecting to DB: %v\n", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db: dbQueries,
	}

	mux := http.NewServeMux()

	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot+"/templates")))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("/", handlerGetIndex)

	// mux.HandleFunc("/api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUsers)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from http://localhost:%s/\n", port)
	log.Fatal(srv.ListenAndServe())
}
