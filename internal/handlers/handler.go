package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-avatar/internal/avatar"
)

// IndexHandler handles the root route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := os.ReadFile("./index.html")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	http.ServeFile(w, r, "./index.html")
}

// AvatarHandler generates and serves the avatar image
func AvatarHandler(w http.ResponseWriter, r *http.Request) {

	// Generate avatar image
	fmt.Println("Generate")
	out, err := avatar.GenerateAvatar()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Image generated!")
	// Serve the image
	w.Header().Set("Content-Type", "image/png")
	w.Write(out.Bytes())
}
