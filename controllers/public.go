package controllers

import (
	"html/template"
	"log"
	"net/http"

	"time.domsay.com/lib"
)

func HomeController(w http.ResponseWriter, r *http.Request) {

	lib.AppRedirect(w, r, "/login", 302)

}

func LoginController(w http.ResponseWriter, r *http.Request) {

	lib.CheckAuth(w, r, true, "/dashboard")

	p := Page{Name: "Login", View: "login", URL: "http://localhost:8080"}

	parsedTemplate, _ := template.ParseFiles("templates/login.html")
	err := parsedTemplate.Execute(w, p)

	if err != nil {
		log.Printf("Error occurred while executing the template  or writing its output")
		return
	}

}
