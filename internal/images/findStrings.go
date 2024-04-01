package images

import (
	"os"
	"path/filepath"
	"strings"
)

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
