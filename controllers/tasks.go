package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"time.domsay.com/lib"
)

type TaskPage struct {
	URL      string
	Name     string
	View     string
	OneData  lib.Task
	AllData  []lib.Task
	Projects map[string]string
}

func TaskHandleTemplate(w http.ResponseWriter, r *http.Request, p TaskPage) {

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

func TasksController(w http.ResponseWriter, r *http.Request) {

	tasksList := lib.GetTasks()

	p := TaskPage{Name: "Tasks", View: "tasks/index", AllData: tasksList}

	TaskHandleTemplate(w, r, p)

}

func TaskAddController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		p := TaskPage{Name: "Add task", View: "tasks/add", Projects: lib.GetProjectsList()}

		log.Println(p)

		TaskHandleTemplate(w, r, p)
	case "POST":

		project_id := r.FormValue("project_id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		start := r.FormValue("start")
		end := r.FormValue("end")

		u := lib.Task{
			ProjectId:   project_id,
			Name:        name,
			Description: description,
			Start:       start,
			End:         end,
			CreatedId:   lib.CookieValid(r),
			ModifiedId:  lib.CookieValid(r),
		}

		log.Println(u)

		lib.InsertTask(u)
		lib.AppRedirect(w, r, "/tasks", 302)

	default:
		// Give an error message.
	}

}

func TaskEditController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		vars := mux.Vars(r)
		id := vars["id"]

		u := lib.GetTask(id)

		log.Println(u)

		p := TaskPage{Name: "Edit task", View: "tasks/edit", OneData: u, Projects: lib.GetProjectsList()}

		TaskHandleTemplate(w, r, p)

	case "POST":

		vars := mux.Vars(r)
		id := vars["id"]

		log.Println(id)

		project_id := r.FormValue("project_id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		start := r.FormValue("start")
		end := r.FormValue("end")

		u := lib.Task{
			ProjectId:   project_id,
			Name:        name,
			Description: description,
			Start:       start,
			End:         end,
			ModifiedId:  lib.CookieValid(r),
		}

		log.Println(u)

		lib.EditTask(id, u)

		lib.AppRedirect(w, r, "/tasks", 302)

	default:
		// Give an error message.
	}

}
