package handlers

import (
	"fmt"
	"go-avatar/internal/avatar"
	"net/http"
	"strings"
)

// AvatarHandler generates and serves the avatar image
func AvatarHandler(w http.ResponseWriter, r *http.Request) {

	// Generate avatar image
	fmt.Println("Generate")
	if fmt.Sprint(r.URL)[:len("/avatar?username=")] != "/avatar?username=" {
		w.Write([]byte("No no no :("))
		return
	}
	api_var := fmt.Sprint(r.URL)[len("/avatar?username="):]
	//fmt.Println("api", api_var)
	strings.Replace(fmt.Sprint(r.URL), api_var, "", 1)
	out, err := avatar.GenerateAvatar(api_var)
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
