package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"shelfgit.com/mdata/metaimage/metaimage"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		path := "./images/" + handler.Filename
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		keys, ReverseColorCounter, TotalPixels := metaimage.ImageProcess(path)
		for _, val := range keys[:5] {
			fmt.Printf("%s %.2f%%\n", ReverseColorCounter[val], ((float64(val) / float64(TotalPixels)) * 100))
			fmt.Fprintf(w, "\n%s %.2f%%\n", ReverseColorCounter[val], ((float64(val) / float64(TotalPixels)) * 100))
		}

	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
