package handlers

import (
	"net/http"
	"time"

	"dev11/calendar"
	"dev11/server/domain"
	"dev11/server/helpers"
)

func response(w http.ResponseWriter, data []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func responseError(w http.ResponseWriter, errStr error, code int) {
	data, err := helpers.SerealizeErr(errStr.Error())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response(w, data, code)
}

func (e *Handler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ev, err := helpers.ParseEvent(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cev := (*calendar.Event)(ev)
	e.cd.CreateEvent(cev)

	data, err := helpers.SerealizeOk(cev.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response(w, data, http.StatusOK)
}

func (e *Handler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ev, err := helpers.ParseEvent(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cev := (*calendar.Event)(ev)
	if err := e.cd.UpdateEvent(cev); err != nil {
		responseError(w, err, http.StatusServiceUnavailable)
		return
	}

	data, err := helpers.SerealizeOk(cev.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response(w, data, http.StatusOK)
}

func (e *Handler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userid, id, err := helpers.ParseDelete(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := e.cd.RemoveEvent(userid, id); err != nil {
		responseError(w, err, http.StatusServiceUnavailable)
		return
	}

	data, err := helpers.SerealizeOk(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response(w, data, http.StatusOK)
}

func (e *Handler) forEventHandler(t time.Duration, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userid, date, err := helpers.ParseUserDate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cev, err := e.cd.GetEvents(userid, date, date.Add(t))
	if err != nil {
		responseError(w, err, http.StatusServiceUnavailable)
		return
	}

	ev := make([]domain.Event, len(cev))
	for k, v := range cev {
		ev[k] = domain.Event(v)
	}

	data, err := helpers.SerealizeEvents(ev)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response(w, data, http.StatusOK)
}

func (e *Handler) ForDayEventHandler(w http.ResponseWriter, r *http.Request) {
	e.forEventHandler(time.Hour*24, w, r)
}

func (e *Handler) ForWeekEventHandler(w http.ResponseWriter, r *http.Request) {
	e.forEventHandler(time.Hour*24*7, w, r)
}

func (e *Handler) ForMonthEventHandler(w http.ResponseWriter, r *http.Request) {
	e.forEventHandler(time.Hour*24*7*31, w, r)
}
