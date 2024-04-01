package main

import (
	"log"
	"net/http"

	"go-avatar/internal/handlers"
	"go-avatar/internal/images"

	"github.com/gorilla/mux"
)

func init() {
	images.Init()
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/avatar", handlers.AvatarHandler).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static"))))

	log.Println("Server started on http://localhost:8055")
	log.Fatal(http.ListenAndServe(":8055", r))
}
