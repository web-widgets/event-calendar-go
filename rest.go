package main

import (
	"net/http"
	"web-widgets/scheduler-go/data"

	"github.com/go-chi/chi"
)

func initRoutes(r chi.Router, dao *data.DAO) {

	r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
		data, err := dao.Events.GetAll()
		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, data)
		}
	})

	r.Post("/events", func(w http.ResponseWriter, r *http.Request) {
		event, err := ParseFormEvent(w, r)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		id, err := dao.Events.Add(event.GetModel())
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		} else {
			format.JSON(w, 200, Response{id})
		}
	})

	r.Put("/events", func(w http.ResponseWriter, r *http.Request) {
		event, err := ParseFormEvent(w, r)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		err = dao.Events.Update(event.GetModel())
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		} else {
			format.JSON(w, 200, nil)
		}
	})

	r.Delete("/events/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")

		err := dao.Events.Delete(id)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		} else {
			format.JSON(w, 200, nil)
		}
	})
}
