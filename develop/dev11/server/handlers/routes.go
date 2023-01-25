package handlers

import (
	"net/http"

	"dev11/calendar"
)

type Handler struct {
	cd  *calendar.Calendar
	mux *http.ServeMux
}

func (e *Handler) createRoutes() {
	e.mux.HandleFunc("/create_event", e.CreateEventHandler)
	e.mux.HandleFunc("/update_event", e.UpdateEventHandler)
	e.mux.HandleFunc("/delete_event", e.DeleteEventHandler)
	e.mux.HandleFunc("/events_for_day", e.ForDayEventHandler)
	e.mux.HandleFunc("/events_for_week", e.ForWeekEventHandler)
	e.mux.HandleFunc("/events_for_month", e.ForMonthEventHandler)
}

func NewRouter(cd *calendar.Calendar) *Handler {
	e := &Handler{
		cd:  cd,
		mux: http.NewServeMux(),
	}

	e.createRoutes()

	return e
}

func (e *Handler) Mux() *http.ServeMux {
	return e.mux
}
