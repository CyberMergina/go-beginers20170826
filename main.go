package main

import (
	"net/http"
	"log"
	"path/filepath"
	"html/template"
	"sync"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})

	if err:=http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal("ListenAndServe",err)
	}
}
