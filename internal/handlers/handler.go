package handlers

import (
	"fmt"
	"go-avatar/internal/avatar"
	"net/http"
)

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
