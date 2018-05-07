package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"time.domsay.com/lib"
)

type SubjectPage struct {
	URL     string
	Name    string
	View    string
	OneData lib.Subject
	AllData []lib.Subject
}

func SubjectHandleTemplate(w http.ResponseWriter, r *http.Request, p SubjectPage) {

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

func SubjectsController(w http.ResponseWriter, r *http.Request) {

	subjectsList := lib.GetSubjects()

	p := SubjectPage{Name: "Subjects", View: "subjects/index", AllData: subjectsList}

	SubjectHandleTemplate(w, r, p)

}

func SubjectAddController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		p := SubjectPage{Name: "Add subject", View: "subjects/add"}

		SubjectHandleTemplate(w, r, p)
	case "POST":

		name := r.FormValue("name")
		vat := r.FormValue("vat")
		address := r.FormValue("address")

		u := lib.Subject{
			Name:       name,
			VAT:        vat,
			Address:    address,
			CreatedId:  lib.CookieValid(r),
			ModifiedId: lib.CookieValid(r),
		}

		log.Println(u)

		lib.InsertSubject(u)
		lib.AppRedirect(w, r, "/subjects", 302)

	default:
		// Give an error message.
	}

}

func SubjectEditController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		vars := mux.Vars(r)
		id := vars["id"]

		u := lib.GetSubject(id)

		log.Println(u)

		p := SubjectPage{Name: "Edit subject", View: "subjects/edit", OneData: u}

		SubjectHandleTemplate(w, r, p)
	case "POST":

		vars := mux.Vars(r)
		id := vars["id"]

		log.Println(id)

		name := r.FormValue("name")
		vat := r.FormValue("vat")
		address := r.FormValue("address")

		u := lib.Subject{
			Name:       name,
			VAT:        vat,
			Address:    address,
			ModifiedId: lib.CookieValid(r),
		}

		log.Println(u)

		lib.EditSubject(id, u)

		lib.AppRedirect(w, r, "/subjects", 302)

	default:
		// Give an error message.
	}

}
