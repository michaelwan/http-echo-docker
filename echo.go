package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"html/template"
)

type RawRequest struct {
	Raw[] byte
}

func handler(w http.ResponseWriter, r *http.Request) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	var templates = template.Must(template.ParseFiles(dir+"/test.html"))

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	rawReq := RawRequest{Raw:dump}

	err = templates.ExecuteTemplate(w, "test.html", rawReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	fmt.Printf("%s", rawReq);

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3322", nil)
}
