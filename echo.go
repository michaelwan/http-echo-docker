package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"html/template"
)

type RawRequest struct {
	Raw[] byte
}

var templates = template.Must(template.ParseFiles("/ADPFedTest.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	rawReq := RawRequest{Raw:dump}

	err = templates.ExecuteTemplate(w, "ADPFedTest.html", rawReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
