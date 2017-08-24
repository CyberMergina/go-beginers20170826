package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte(`Hello, World`))
	})
	if err:=http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal("ListenAndServe",err)
	}
}