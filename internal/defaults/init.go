package defaults

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {

	flush.TmpFolder = "./output/tmp/"
	firstRun, err := os.ReadFile(".hasRun")
	if err != nil {
		if fmt.Sprint(err) == "open .hasRun: no such file or directory" {
			os.WriteFile(".hasRun", []byte("false"), 0777)
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
			fmt.Println("Downloading images..")
			tmp.DownloadImage(url, "")
		}
		fmt.Println("Removing temporary files..")
		MoveDir("output/tmp/minipix/clothing", "output/files")
		os.RemoveAll("output/tmp")
		os.WriteFile(".hasRun", []byte("true"), 0755)
	}
}
