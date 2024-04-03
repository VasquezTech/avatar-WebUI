package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"go-avatar/internal/handlers"
	"go-avatar/internal/images"

	"github.com/gorilla/mux"
)

//go:embed static/dist
var app embed.FS

func main() {

	images.Init()
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	r := mux.NewRouter()

	dist, err := fs.Sub(app, "static/dist")
	if err != nil {
		log.Fatalf("sub error")
		return
	}
	r.Handle("/", http.FileServer(http.FS(dist)))
	//	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/avatar", handlers.AvatarHandler).Methods("GET")

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static"))))

	log.Println("Server started on http://localhost:8055")
	log.Fatal(http.ListenAndServe(":8055", r))
}
