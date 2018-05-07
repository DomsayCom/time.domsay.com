package controllers

import (
	"html/template"
	"log"
	"net/http"

	"time.domsay.com/lib"
)

type Page struct {
	URL  string
	Name string
	View string
	Data map[string]string
}

func HandleTemplate(w http.ResponseWriter, r *http.Request, p Page) {

	lib.CheckAuth(w, r, false, "/login")

	parsedTemplate, _ := template.ParseFiles(
		"templates/index.html",
		"templates/sidebar.html",
		"templates/headerbar.html",
		"views/"+p.View+".html",
	)

	p.URL = "http://localhost:8080"

	err := parsedTemplate.Execute(w, p)

	if err != nil {
		log.Printf("Error occurred while executing the template  or writing its output")
		return
	}

}
