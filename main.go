package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot+"/templates")))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("/", handlerGetIndex)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from http://localhost:%s/\n", port)
	log.Fatal(srv.ListenAndServe())
}
