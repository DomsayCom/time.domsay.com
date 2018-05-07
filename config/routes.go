package config

import (
	"net/http"

	"time.domsay.com/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
	Route{"home", "GET", "/", controllers.HomeController},

	Route{"login", "GET", "/login", controllers.LoginController},

	Route{"user_login", "POST", "/user/login", controllers.UserLoginController},
	Route{"user_logout", "GET", "/user/logout", controllers.UserLogoutController},

	Route{"dashboard", "GET", "/dashboard", controllers.DashboardController},

	Route{"users", "GET", "/users", controllers.UsersController},

	Route{"user_add", "GET", "/users/add", controllers.UserAddController},
	Route{"user_add", "POST", "/users/add", controllers.UserAddController},

	Route{"user_edit", "GET", "/users/edit/{id}", controllers.UserEditController},
	Route{"user_edit", "POST", "/users/edit/{id}", controllers.UserEditController},

	Route{"subjects", "GET", "/subjects", controllers.SubjectsController},

	Route{"subjects_add", "GET", "/subjects/add", controllers.SubjectAddController},
	Route{"subjects_add", "POST", "/subjects/add", controllers.SubjectAddController},

	Route{"subjects_edit", "GET", "/subjects/edit/{id}", controllers.SubjectEditController},
	Route{"subjects_edit", "POST", "/subjects/edit/{id}", controllers.SubjectEditController},

	Route{"projects", "GET", "/projects", controllers.ProjectsController},

	Route{"projects_add", "GET", "/projects/add", controllers.ProjectAddController},
	Route{"projects_add", "POST", "/projects/add", controllers.ProjectAddController},

	Route{"projects_edit", "GET", "/projects/edit/{id}", controllers.ProjectEditController},
	Route{"projects_edit", "POST", "/projects/edit/{id}", controllers.ProjectEditController},

	Route{"tasks", "GET", "/tasks", controllers.TasksController},

	Route{"tasks_add", "GET", "/tasks/add", controllers.TaskAddController},
	Route{"tasks_add", "POST", "/tasks/add", controllers.TaskAddController},

	Route{"tasks_edit", "GET", "/tasks/edit/{id}", controllers.TaskEditController},
	Route{"tasks_edit", "POST", "/tasks/edit/{id}", controllers.TaskEditController},
}
