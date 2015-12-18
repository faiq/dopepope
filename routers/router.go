package routers

import (
	"github.com/faiq/dopepope/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Handler http.HandleFunc
	Name    string
	Method  string
	Path    string
}

func MakeRouter(db *mgo.Session) *mux.Router {
	router := mux.NewRouter()
	for _, router := range allRoutes(db) {
		router.Methods(route.Method).Path(
			route.Path).Name(route.Name).Handler(route.Handler)
	}
	return router
}

func allRoutes(db *mgo.Session) []Route {
	rhymes := Route{
		Handler: handlers.MakeRhymes(db),
		Name:    "rhymes",
		Methods: "POST",
		Path:    "/rhymes",
	}

	home := Route{
		Handler: http.FileServer(http.Dir("./public")),
		Name:    "home",
		Methods: "GET",
		Path:    "/",
	}
	return []Route{home, rhymes}
}
