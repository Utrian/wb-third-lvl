package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wb-third-lvl/develop/dev11/internal/calendar"

	"github.com/sirupsen/logrus"
)

type router struct {
	storage *calendar.Calendar
}

func NewRouter() *router {
	return &router{storage: calendar.NewCalendar()}
}

func (h *router) Register(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", h.createEventHandler)
	mux.HandleFunc("/update_event", h.updateEventHandler)
	mux.HandleFunc("/delete_event", h.deleteEventHandler)
	mux.HandleFunc("/events_for_day", h.eventsForDayHandler)
	mux.HandleFunc("/events_for_week", h.eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", h.eventsForMonthHandler)
}

func (h *router) createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		logrus.Error(errors.New("invalid method"))
		return
	}

	query := r.URL.Query()
	evn, err := h.storage.NewEvent(&query)
	if err != nil {
		w.WriteHeader(400)
		logrus.Error(err)
		return
	}

	bytes, err := json.Marshal(evn)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *router) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		logrus.Error(errors.New("invalid method"))
		return
	}

	query := r.URL.Query()
	evn, err := h.storage.Update(&query)
	if err != nil {
		w.WriteHeader(400)
		logrus.Error(err)
		return
	}

	bytes, err := json.Marshal(evn)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *router) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		logrus.Error(errors.New("invalid method"))
		return
	}

	query := r.URL.Query()
	evn, err := h.storage.Delete(&query)
	if err != nil {
		w.WriteHeader(400)
		logrus.Error(err)
		return
	}

	bytes, err := json.Marshal(evn)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *router) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		logrus.Error(errors.New("invalid method"))
		return
	}

	query := r.URL.Query()
	evns, err := h.storage.DayFilter(&query)
	if err != nil {
		w.WriteHeader(400)
		logrus.Error(err)
		return
	}

	bytes, err := json.Marshal(evns)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *router) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		logrus.Error(errors.New("invalid method"))
		return
	}

	query := r.URL.Query()
	evns, err := h.storage.WeekFilter(&query)
	if err != nil {
		w.WriteHeader(400)
		logrus.Error(err)
		return
	}

	bytes, err := json.Marshal(evns)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *router) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		logrus.Error(errors.New("invalid method"))
		return
	}

	query := r.URL.Query()
	evns, err := h.storage.MonthFilter(&query)
	if err != nil {
		w.WriteHeader(400)
		logrus.Error(err)
		return
	}

	bytes, err := json.Marshal(evns)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
