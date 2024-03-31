package main

import (
    "html/template"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/avatarme/identicon"
)

// IndexHandler handles the root route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, nil)
}

// AvatarHandler generates and serves the avatar image
func AvatarHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    text := r.Form.Get("text")

    // Generate avatar image based on input text
    img := identicon.Render(text, 200)

    // Serve the image
    w.Header().Set("Content-Type", "image/png")
    w.Write(img)
}

func main() {
    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/avatar", AvatarHandler).Methods("POST")
    // Static
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Start
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
