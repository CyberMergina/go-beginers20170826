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

type Product struct {
	Name	string
}

type Data struct {
	Products	[]*Product
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	p1 := Product{Name:"りんご"}
	p2 := Product{Name:"なし"}
	p3 := Product{Name:"ばなな"}
	d := Data{ Products: []*Product{&p1, &p2, &p3} }
	t.templ.Execute(w, d)

}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})

	if err:=http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal("ListenAndServe",err)
	}
}
