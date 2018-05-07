package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"time.domsay.com/lib"
)

type ProjectPage struct {
	URL      string
	Name     string
	View     string
	OneData  lib.Project
	AllData  []lib.Project
	Subjects map[string]string
}

func ProjectHandleTemplate(w http.ResponseWriter, r *http.Request, p ProjectPage) {

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

func ProjectsController(w http.ResponseWriter, r *http.Request) {

	projectsList := lib.GetProjects()

	p := ProjectPage{Name: "Projects", View: "projects/index", AllData: projectsList}

	ProjectHandleTemplate(w, r, p)

}

func ProjectAddController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		p := ProjectPage{Name: "Add project", View: "projects/add", Subjects: lib.GetSubjectsList()}

		log.Println(p)

		ProjectHandleTemplate(w, r, p)
	case "POST":

		subject_id := r.FormValue("subject_id")
		name := r.FormValue("name")

		u := lib.Project{
			SubjectId:  subject_id,
			Name:       name,
			CreatedId:  lib.CookieValid(r),
			ModifiedId: lib.CookieValid(r),
		}

		log.Println(u)

		lib.InsertProject(u)
		lib.AppRedirect(w, r, "/projects", 302)

	default:
		// Give an error message.
	}

}

func ProjectEditController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		vars := mux.Vars(r)
		id := vars["id"]

		u := lib.GetProject(id)

		log.Println(u)

		p := ProjectPage{Name: "Edit project", View: "projects/edit", OneData: u, Subjects: lib.GetSubjectsList()}

		ProjectHandleTemplate(w, r, p)

	case "POST":

		vars := mux.Vars(r)
		id := vars["id"]

		log.Println(id)

		subject_id := r.FormValue("subject_id")
		name := r.FormValue("name")

		u := lib.Project{
			SubjectId:  subject_id,
			Name:       name,
			ModifiedId: lib.CookieValid(r),
		}

		log.Println(u)

		lib.EditProject(id, u)

		lib.AppRedirect(w, r, "/projects", 302)

	default:
		// Give an error message.
	}

}
