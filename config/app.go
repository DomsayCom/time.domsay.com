package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func AddRoutes(router *mux.Router) *mux.Router {

	for _, route := range RoutesList {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router

}

func StartApp() {

	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	// add static file prefix
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	// demo
	//AddRoutes(muxRouter.PathPrefix("/demo").Subrouter())

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

}
