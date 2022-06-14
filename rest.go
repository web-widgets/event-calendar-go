package main

import (
	"net/http"
	"web-widgets/scheduler-go/data"

	"github.com/go-chi/chi"
)

func initRoutes(r chi.Router, dao *data.DAO) {

	r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
		data, err := dao.Events.GetAll()
		sendResponse(w, data, err)
	})

	r.Post("/events", func(w http.ResponseWriter, r *http.Request) {
		event := data.EventUpdate{}
		err := ParseForm(w, r, &event)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		id, err := dao.Events.Add(&event)
		sendResponse(w, Response{id}, err)
	})

	r.Put("/events/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		event := data.EventUpdate{}
		err := ParseForm(w, r, &event)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		err = dao.Events.Update(id, &event)
		sendResponse(w, nil, err)
	})

	r.Delete("/events/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		err := dao.Events.Delete(id)
		sendResponse(w, nil, err)
	})

	r.Get("/calendars", func(w http.ResponseWriter, r *http.Request) {
		data, err := dao.Calendars.GetAll()
		sendResponse(w, data, err)
	})

	r.Post("/calendars", func(w http.ResponseWriter, r *http.Request) {
		calendar := data.CalendarUpdate{}
		err := ParseForm(w, r, &calendar)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		id, err := dao.Calendars.Add(&calendar)
		sendResponse(w, Response{id}, err)
	})

	r.Put("/calendars/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		calendar := data.CalendarUpdate{}
		err := ParseForm(w, r, &calendar)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		err = dao.Calendars.Update(id, &calendar)
		sendResponse(w, nil, err)
	})

	r.Delete("/calendars/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		err := dao.Calendars.Delete(id)
		sendResponse(w, nil, err)
	})
}

func sendResponse(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		format.Text(w, 500, err.Error())
	} else {
		format.JSON(w, 200, data)
	}
}
