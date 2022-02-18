package main

import (
	"github.com/gorilla/mux"
	"github.com/mustamohamed/weekplanner/api"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	api.RegisterDefaultPlanHandler(router)
	http.Handle("/", router)
	http.ListenAndServe(":5000", router)
}
