package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
