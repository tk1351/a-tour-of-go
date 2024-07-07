package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}
