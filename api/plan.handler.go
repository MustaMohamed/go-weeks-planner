package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mustamohamed/weekplanner/errors"
	"github.com/mustamohamed/weekplanner/models"
	"github.com/mustamohamed/weekplanner/services"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PlanHandler struct {
	service services.IPlanService
}

func NewPlanHandler(db *gorm.DB, ctx context.Context) *PlanHandler {
	return &PlanHandler{
		service: services.NewPlanService(db, ctx),
	}
}

func (ph *PlanHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	plans, err := ph.service.GetAll()
	if err != nil {
		if _, ok := err.(*errors.NotFoundRecordError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	resp, err := json.Marshal(plans)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (ph *PlanHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	plan, err := ph.service.GetById(uint(id))
	if err != nil {
		if _, ok := err.(*errors.NotFoundRecordError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	resp, err := json.Marshal(plan)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (ph *PlanHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	println(string(body))
	plan := models.Plan{}
	err := json.Unmarshal(body, &plan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	savedPlan, err := ph.service.Create(plan)
	if err != nil {
		if _, ok := err.(*errors.NotFoundRecordError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(savedPlan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (ph *PlanHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	plan, err := ph.service.GetById(uint(id))
	if err != nil {
		if _, ok := err.(*errors.NotFoundRecordError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	println(string(body))
	newPlan := &models.Plan{}
	err = json.Unmarshal(body, plan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newPlan.ID = plan.ID
	savedPlan, err := ph.service.Update(*plan)
	if err != nil {
		if _, ok := err.(*errors.NotFoundRecordError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(savedPlan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (ph *PlanHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = ph.service.DeleteById(uint(id))
	if err != nil {
		if _, ok := err.(*errors.NotFoundRecordError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
