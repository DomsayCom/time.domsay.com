package controllers

import (
	"html/template"
	"log"
	"net/http"

	"time.domsay.com/lib"
)

type DashboardPage struct {
	URL        string
	Name       string
	View       string
	AppCounter lib.Counter
	Projects   map[string]lib.ProjectDetails
}

func DashboardHandleTemplate(w http.ResponseWriter, r *http.Request, p DashboardPage) {

	lib.CheckAuth(w, r, false, "/login")

	p.URL = "http://localhost:8080"

	parsedTemplate, _ := template.ParseFiles(
		"templates/index.html",
		"templates/sidebar.html",
		"templates/headerbar.html",
		"views/"+p.View+".html",
	)
	err := parsedTemplate.Execute(w, p)

	if err != nil {
		log.Printf("Error occurred while executing the template  or writing its output")
		return
	}

}

func DashboardController(w http.ResponseWriter, r *http.Request) {

	projects := lib.GetProjectsDetails()

	p := DashboardPage{Name: "Dashboard", View: "dashboard/index", AppCounter: lib.GetCounter(), Projects: projects}

	DashboardHandleTemplate(w, r, p)

}
