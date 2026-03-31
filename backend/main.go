package main

import (
	"log"
	"net/http"
	"path/filepath"
	"tcc/backend/app/httpapi"
	"tcc/backend/app/store"
)

const serverAddress = ":8080"

func main() {
	commentStore, err := store.NewCommentStore(filepath.Join("data", "comments.json"))
	if err != nil {
		log.Fatalf("init store: %v", err)
	}

	handler := httpapi.NewHandler(commentStore)

	log.Printf("TCC backend listening on %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, httpapi.WithCORS(handler.Routes())); err != nil {
		log.Fatal(err)
	}
}
