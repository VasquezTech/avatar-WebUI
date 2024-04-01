package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	// opts := debugR(w, r)
	// url := fmt.Sprintf("https://funky-pixel-avatars.p.rapidapi.com/api/v1/avatar/generate/user?g=%s&uname=%s&fe=%s", opts.Gender, opts.Username, opts.ImgType)
	url := ""
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "c836c57203msh667ac0b1c13cc91p13739bjsn1180f78596c6")
	req.Header.Add("X-RapidAPI-Host", "funky-pixel-avatars.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	w.Header().Set("Content-Type", "image/png")
	w.Write(body)

}

func debugR(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.FormValue("username"))

}
