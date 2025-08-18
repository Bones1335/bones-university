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

	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot+"/app/.")))
	mux.Handle("/app/", fsHandler)

	// Admin endpoints
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /admin/roles", apiCfg.handlerCreateRoles)
	mux.HandleFunc("PUT /admin/users_roles/{user_id}", apiCfg.handlerUpdateUsersRole)
	mux.HandleFunc("POST /admin/degrees", apiCfg.handlerCreateDegrees)
	mux.HandleFunc("POST /admin/years", apiCfg.handlerCreateYear)

	// User endpoints
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUsers)
	mux.HandleFunc("GET /api/users/{user_id}", apiCfg.handlerGetUser)
	mux.HandleFunc("PUT /api/users/{user_id}", apiCfg.handlerUpdateUsers)

	// TODO: Degree endpoints
	mux.HandleFunc("GET /api/degrees", apiCfg.handlerGetDegrees)

	// TODO: Students_programs endpoints
	mux.HandleFunc("POST /api/students_programs", apiCfg.handlerCreateStudentsProgram)

	// TODO: Course endpoints
	mux.HandleFunc("POST /api/courses", apiCfg.handlerCreateCourses)

	// TODO: Degrees_Courses endpoints
	mux.HandleFunc("POST /api/degrees_courses", apiCfg.handlerCreateDegreesCourses)

	// TODO: Assignment endpoints
	mux.HandleFunc("POST /api/assignments", apiCfg.handlerCreateAssignments)

	// TODO: Course_Assignment endpoints

	// TODO: Grades ednpoints

	// TODO: Assignment_Grade endpoints

	// TODO: Internship enspoints

	// TODO: Internship_User(student) endpoints

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from http://localhost:%s/app/\n", port)
	log.Fatal(srv.ListenAndServe())
}
