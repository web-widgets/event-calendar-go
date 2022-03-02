package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"web-widgets/scheduler-go/data"

	"github.com/go-chi/chi"
	remote "github.com/mkozhukh/go-remote"
)

func initRoutes(r chi.Router, dao *data.DAO, hub *remote.Hub) {

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

	// DEMO ONLY, imitate login
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		uid, _ := strconv.Atoi(r.URL.Query().Get("id"))
		device := newDeviceID()
		token, err := createUserToken(uid, device)
		if err != nil {
			log.Println("[token]", err.Error())
		}
		w.Write(token)
	})
}

var dID int

func init() {
	dID = int(time.Now().Unix())
}

func newDeviceID() int {
	dID += 1
	return dID
}
