package main

import (
	"net/http"
	"github.com/russross/blackfriday"
	"os"
)

func main() {

puerto := os.Getenv("PORT")

if puerto == "" {
		puerto = "8080"
	}
	http.HandleFunc("/markdown", GeneraDesdeMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+puerto, nil)

}

func GeneraDesdeMarkdown(rw http.ResponseWriter, r *http.Request) {
	html := blackfriday.MarkdownCommon([]byte (r.FormValue("cuerpo")))
	rw.Write(html)
}

