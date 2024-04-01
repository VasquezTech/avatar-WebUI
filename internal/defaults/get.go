package defaults

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
	Current_html string
	Pwds         []string
	TmpFolders   []string
	Urls         []string
}

var flush Files

// ex first string:  "a href=\"/minipix/"
// ex second string: "href=\""
// ex thirds string: "\""
func (F Files) Set_urls(firstString, secondString, thirdString string) Files {

	lines := strings.Split(F.Current_html, "\n")

	for i := range lines {
		if lines[i] != "" {
			if strings.Contains(lines[i], firstString) {
				// fmt.Println("line:", lines[i], "\nSecond:", secondString)
				urlStr1 := strings.Split(lines[i], secondString)[1]
				url := strings.Split(urlStr1, thirdString)[0]

				F.Urls = append(F.Urls, url)
				// fmt.Println("New url:", url)
			}
		}
	}

	return F
}
func (F Files) Set_html_folders(firstString, secondString, thirdString string) Files {

	lines := strings.Split(F.Current_html, "\n")

	for i := range lines {
		if strings.Contains(lines[i], firstString) {

			urlStr1 := strings.Split(lines[i], secondString)[1]
			url := strings.Split(urlStr1, thirdString)[0]

			F.TmpFolders = append(F.TmpFolders, url)
			// fmt.Println("New folder:", url)
		}
	}

	return F
}
func (F Files) Set_pwds() Files {

	for _, pwd := range F.Urls {
		pwds := strings.Split(pwd, "/")
		if len(pwds) < 4 {
			continue
		}
		var folder, file string
		if strings.Contains(pwd, ".") {
			file = pwds[len(pwds)-2] + "_" + pwds[len(pwds)-1]
			for i := range pwds {
				if pwds[i] != pwds[len(pwds)-1] && pwds[i] != pwds[len(pwds)-2] && pwds[i] != "" {
					folder = folder + "/" + pwds[i]
					// fmt.Println("New folder:", folder)
				}
			}
		} else {
			for i := range pwds {
				if pwds[i] != pwds[len(pwds)-1] && pwds[i] != "" {
					folder = folder + "/" + pwds[i]

					os.Mkdir(filepath.Join(F.TmpFolder, folder), 0777)

					// fmt.Println("New folder:", folder)
				}
			}
		}

		fullpwd := filepath.Join(folder, file) // Save
		newFolder := true
		for i := range F.TmpFolders {
			if F.TmpFolders[i] == folder {
				newFolder = false
			}
		}
		if newFolder {
			F.TmpFolders = append(F.TmpFolders, folder)
		}
		F.Pwds = append(F.Pwds, fullpwd)
	}
	return F
}

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
			if len(F.TmpFolders) > 0 {

				// fmt.Println("file", true)
				file = F.TmpFolder + u + ext
			}
		}
		// fmt.Println("file", file)
		if strings.Contains(file, ".png") {

			// fmt.Println("file", false)
			file = F.TmpFolder + "minipix/clothing" + file
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
			if len(F.TmpFolders) > 0 {

				// fmt.Println("file", true)
				file = F.TmpFolder + u + ext
			}
		}
		// fmt.Println("file", file)
		if strings.Contains(file, ".png") {

			// fmt.Println("file", false)
			file = F.TmpFolder + "minipix/clothing" + file
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
