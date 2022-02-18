package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mustamohamed/weekplanner/database"
	"net/http"
)

const BaseRoute = "/api"

func RegisterPlanHandler(router *mux.Router, ph *PlanHandler) {
	router.HandleFunc("/plans", ph.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/plans", ph.Create).Methods(http.MethodPost)
	router.HandleFunc("/plans/{id}", ph.GetById).Methods(http.MethodGet)
	router.HandleFunc("/plans/{id}", ph.UpdateById).Methods(http.MethodPut)
	router.HandleFunc("/plans/{id}", ph.DeleteById).Methods(http.MethodDelete)
}

func RegisterDefaultPlanHandler(router *mux.Router) {
	ctx := context.Background()
	defer ctx.Done()
	RegisterPlanHandler(router, NewPlanHandler(database.GetConnection(), ctx))
}
