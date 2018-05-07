package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"time.domsay.com/lib"
)

type UsersPage struct {
	URL     string
	Name    string
	View    string
	OneData lib.User
	AllData []lib.User
}

func UsersHandleTemplate(w http.ResponseWriter, r *http.Request, p UsersPage) {

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

func UsersController(w http.ResponseWriter, r *http.Request) {

	usersList := lib.GetUsers()

	p := UsersPage{Name: "Users", View: "users/index", AllData: usersList}

	UsersHandleTemplate(w, r, p)

}

func UserAddController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		p := UsersPage{Name: "Add user", View: "users/add"}

		UsersHandleTemplate(w, r, p)
	case "POST":

		username := r.FormValue("username")
		password := r.FormValue("password")
		active := r.FormValue("active")

		a, _ := strconv.Atoi(active)

		u := lib.User{
			Username:   username,
			Password:   password,
			Active:     a,
			CreatedId:  lib.CookieValid(r),
			ModifiedId: lib.CookieValid(r),
		}

		log.Println(u)

		lib.InsertUser(u)
		lib.AppRedirect(w, r, "/users", 302)

	default:
		// Give an error message.
	}

}

func UserEditController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		vars := mux.Vars(r)
		id := vars["id"]

		u := lib.GetUser(id)

		log.Println(u)

		p := UsersPage{Name: "Edit user", View: "users/edit", OneData: u}

		UsersHandleTemplate(w, r, p)
	case "POST":

		vars := mux.Vars(r)
		id := vars["id"]

		log.Println(id)

		username := r.FormValue("username")
		active := r.FormValue("active")

		a, _ := strconv.Atoi(active)

		u := lib.User{
			Username:   username,
			Active:     a,
			ModifiedId: lib.CookieValid(r),
		}

		log.Println(u)

		lib.EditUser(id, u)

		lib.AppRedirect(w, r, "/users", 302)

	default:
		// Give an error message.
	}

}

func UserLoginController(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")
	target := "/login"

	if username != "" && password != "" {

		uid := lib.CheckUser(username, password)
		if uid != "" {

			lib.SetSession(uid, w)
			target = "/dashboard"

		}

	}

	lib.AppRedirect(w, r, target, 302)

}

func UserLogoutController(w http.ResponseWriter, r *http.Request) {

	lib.ClearSession(w)
	lib.AppRedirect(w, r, "/login", 302)

}
