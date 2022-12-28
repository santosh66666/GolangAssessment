package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func download(url string) (err error) {
	filename := path.Base(url)
	allowedExtensions := []string{".svg",
	".jpg",
	".png",
	".jpeg"}
	fileext := filepath.Ext(url)
	validExt:=contains(allowedExtensions, fileext)
	if validExt == false{
		fmt.Println("This ", fileext, "extension is not supported")
		return
	}
	fmt.Println("fileext ", fileext)
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
			return
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
			return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func main() {
	download("https://www.gstatic.com/webp/gallery3/1.sm.png")
}