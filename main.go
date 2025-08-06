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
	db        *database.Queries
	platform  string
	jwtSecret string
}

func main() {
	env.SetEnv(".env")

	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	platform := os.Getenv("PLATFORM")

	const filepathRoot = "."
	const port = "8080"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Error connecting to DB: %v\n", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db:        dbQueries,
		platform:  platform,
		jwtSecret: jwtSecret,
	}

	mux := http.NewServeMux()

	fsHandler := http.StripPrefix("/static", http.FileServer(http.Dir(filepathRoot+"/static")))
	mux.Handle("/static/", fsHandler)

	// Page Views
	mux.HandleFunc("/", handlerGetIndex)
	mux.HandleFunc("/enrollment", handlerCreateEnrollment)
	mux.HandleFunc("/login", handlerLogin)
	mux.HandleFunc("/student_dashboard", handlerStudentDashboard)
	mux.HandleFunc("/admin_dashboard", handlerAdminDashboard)

	// API endpoints
	mux.HandleFunc("POST /api/reset", apiCfg.handlerReset)
	mux.HandleFunc("/api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUsers)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from http://localhost:%s/\n", port)
	log.Fatal(srv.ListenAndServe())
}
