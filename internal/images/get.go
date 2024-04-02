package images

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// https://www.avatarsinpixels.com/
// a href="/minipix/
// img src="/minipix/Body/1/thumbnail.png"
func Get(url string) string {
	// fmt.Println("Downloading", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed toget: %v", err)
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}

type Files struct {
	TmpFolder    string
	StaticFiles  string
	Current_html string
	Pwds         []string
	TmpFolders   []string
	Urls         []string
}

var flush Files

func (F Files) DownloadFile(url, ext string) Files {
	URL := url
	for _, u := range F.TmpFolders {
		// fmt.Println("GetURL", u)
		url = URL + u
		var f []byte

		f = []byte(Get("https://" + url))
		var file, folder string

		if strings.Contains(url, ".png") {
			pwds := strings.Split(url, "/")
			// fmt.Println(url)
			file = pwds[len(pwds)-2] + "_" + pwds[len(pwds)-1]
			for i := range pwds {
				if i > 1 && pwds[i] != pwds[len(pwds)-1] && pwds[i] != pwds[len(pwds)-2] && pwds[i] != "" {
					folder = folder + "/" + pwds[i]
					// fmt.Println("New folder:", folder)
				}
			}

			file = filepath.Join(folder, file)
		} else {
			if len(F.StaticFiles) > 0 {

				// fmt.Println("file", true)
				file = F.StaticFiles + u + ext
			}
		}
		// fmt.Println("file", file)
		if strings.Contains(file, ".png") {

			// fmt.Println("file", false)
			file = F.StaticFiles + "minipix/clothing" + file
		}
		// fmt.Println("Saving to:", file)
		_, err := os.Create(file)
		if err != nil {
			continue
		}
		err = os.WriteFile(file, f, 0755)
		if err != nil {
			continue
		}
	}
	return F
}

// CLears the struct
func (F Files) Flush() Files {
	F = flush
	return F
}

func (F Files) DownloadImage(url, ext string) Files {
	URL := url
	for _, u := range F.Urls {
		// fmt.Println("GetURL", u)
		url = URL + u

		f := []byte(Get("https://" + url))
		var file, folder string

		if strings.Contains(url, ".png") {
			pwds := strings.Split(url, "/")
			// fmt.Println(url)
			file = pwds[len(pwds)-2] + "_" + pwds[len(pwds)-1]
			for i := range pwds {
				if i > 1 && pwds[i] != pwds[len(pwds)-1] && pwds[i] != pwds[len(pwds)-2] && pwds[i] != "" {
					folder = folder + "/" + pwds[i]
					// fmt.Println("New folder:", folder)
				}
			}

			file = filepath.Join(folder, file)
		} else {
			if len(F.TmpFolders) > 0 {

				// fmt.Println("file", true)
				file = F.StaticFiles + u + ext
			}
		}
		// fmt.Println("file", file)
		if strings.Contains(file, ".png") {

			// fmt.Println("file", false)
			file = F.StaticFiles + "minipix/clothing" + file
		}
		// fmt.Println("Saving to:", file)
		_, err := os.Create(file)
		if err != nil {
			continue
		}
		err = os.WriteFile(file, f, 0755)
		if err != nil {
			continue
		}
	}
	return F
}
