package images

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {

	flush.TmpFolder = "./tmp/"
	flush.StaticFiles = flush.TmpFolder // "./static/dist/files"
	_, err := os.Stat(flush.StaticFiles)
	if os.IsNotExist(err) {
		os.MkdirAll(flush.StaticFiles, 0755)
	}
	firstRun, err := os.ReadFile("hasRun")
	if err != nil {
		if fmt.Sprint(err) == "open hasRun: no such file or directory" {
			os.WriteFile("hasRun", []byte("false"), 0777)
			Init()
			return
		}
	}
	flush.Current_html = ""
	flush.Pwds = []string{}
	flush.Urls = []string{}
	tmp := flush
	if string(firstRun) == "false" {
		fmt.Println("Gathering first time files. This might take a minute or two..")
		err = os.Mkdir(tmp.TmpFolder, 0777)
		if err != nil {
			log.Println(err)
		}
		tmp.Current_html = Get("https://www.avatarsinpixels.com/minipix/clothing/Body")
		fmt.Println("Setting urls")
		tmp = tmp.Set_urls("a href=\"/minipix/", "href=\"", "\"")
		fmt.Println("Setting pwds")

		tmp = tmp.Set_html_folders("a href=\"/minipix/", "href=\"", "\"")
		for _, folder := range tmp.TmpFolders {
			err := os.MkdirAll(tmp.TmpFolder+folder, 0777)
			if err != nil {
				log.Println(err)
			}

		}
		//tmp = tmp.Set_pwds()
		url := "www.avatarsinpixels.com"

		fmt.Println("Downloading html files")
		tmp.DownloadFile(url, ".html")
		fmt.Println("Downloading images")
		urls := tmp.Urls
		for i := range urls {
			fmt.Println("Checking", filepath.Join(tmp.TmpFolder, urls[i]+".html"))

			f, err := os.ReadFile(filepath.Join(tmp.TmpFolder, urls[i]+".html"))
			if err != nil {
				log.Println(err)
			}
			tmp = tmp.Flush()

			tmp.Current_html = string(f)
			tmp.Urls = append(tmp.Urls, urls[i])
			tmp = tmp.Set_urls("img src=\"/minipix", "src=\"", "\"")
			tmp = tmp.Set_pwds()
			tmp.DownloadImage(url, "")
		}
		fmt.Println("Removing temporary files..")
		MoveDir("./tmp/minipix/clothing", tmp.StaticFiles)

		os.RemoveAll("./tmp")
		os.WriteFile("hasRun", []byte("true"), 0755)

		// Try to eclude any body from traits
		folders, _ := os.ReadDir("./files")
		for _, folder := range folders {
			if !folder.IsDir() || folder.Name() == "Body" {
				continue
			}
			files, _ := os.ReadDir("./files/" + folder.Name()) //Body/0_thumbnail.png")
			for _, file := range files {
				p1 := "./files/Body/0_thumbnail.png"
				p2 := filepath.Join("./files", folder.Name(), file.Name())
				CleanupIMG(p1, p2)
			}
		}
	}
}
