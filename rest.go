package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"web-widgets/scheduler-go/api"
	"web-widgets/scheduler-go/data"

	"github.com/go-chi/chi"
	go_remote "github.com/mkozhukh/go-remote"
)

var dID int

func init() {
	dID = int(time.Now().Unix())
}

func initRoutes(r chi.Router, dao *data.DAO, hub *go_remote.Hub) {

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
		if sendResponse(w, Response{id}, err) {
			e, _ := dao.Events.GetOne(id)
			hub.Publish("events", api.EventItem{
				Type:  "add-event",
				From:  getDeviceID(r),
				Event: &e,
			})
		}
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
		if sendResponse(w, nil, err) {
			e, _ := dao.Events.GetOne(id)
			hub.Publish("events", api.EventItem{
				Type:  "update-event",
				From:  getDeviceID(r),
				Event: &e,
			})
		}
	})

	r.Delete("/events/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		err := dao.Events.Delete(id)
		if sendResponse(w, nil, err) {
			hub.Publish("events", api.EventItem{
				Type:  "delete-event",
				From:  getDeviceID(r),
				Event: &data.Event{ID: id},
			})
		}
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
		if sendResponse(w, Response{id}, err) {
			c, _ := dao.Calendars.GetOne(id)
			hub.Publish("calendars", api.EventCalendar{
				Type:     "add-calendar",
				From:     getDeviceID(r),
				Calendar: c,
			})
		}
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
		if sendResponse(w, nil, err) {
			c, _ := dao.Calendars.GetOne(id)
			hub.Publish("calendars", api.EventCalendar{
				Type:     "update-calendar",
				From:     getDeviceID(r),
				Calendar: c,
			})
		}
	})

	r.Delete("/calendars/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		err := dao.Calendars.Delete(id)
		if sendResponse(w, nil, err) {
			hub.Publish("calendars", api.EventCalendar{
				Type:     "delete-calendar",
				From:     getDeviceID(r),
				Calendar: &data.Calendar{ID: id},
			})
		}

		sendResponse(w, Response{id}, err)
	})

	r.Get("/uploads/{id}/{name}", func(w http.ResponseWriter, r *http.Request) {
		res, err := dao.Files.ToResponse(w, NumberParam(r, "id"))

		if err != nil {
			format.Text(w, 500, err.Error())
		} else if !res {
			format.Text(w, 500, "")
		}
	})

	r.Post("/uploads", func(w http.ResponseWriter, r *http.Request) {
		rec, err := dao.Files.FromRequest(r, "upload")
		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, rec)
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

func sendResponse(w http.ResponseWriter, data interface{}, err error) bool {
	if err != nil {
		format.Text(w, 500, err.Error())
	} else {
		format.JSON(w, 200, data)
	}
	return err == nil
}

func newDeviceID() int {
	dID += 1
	return dID
}
