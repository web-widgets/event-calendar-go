package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type EventDemo struct {
	ID int `json:"id"`
	EventUpdate
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func dataDown(d *DAO) {
	d.mustExec("DELETE from events")
	d.mustExec("DELETE from calendars")
	d.mustExec("DELETE from binary_data")
}

func dataUp(d *DAO) (err error) {
	tempEvents := make([]EventDemo, 0)
	err = parseDemodata(&tempEvents, "./demodata/events.json")
	if err != nil {
		log.Fatal(err)
	}
	events := make([]Event, len(tempEvents))
	for i := range tempEvents {
		events[i] = *tempEvents[i].GetModel()
	}

	calendars := make([]Calendar, 0)
	err = parseDemodata(&calendars, "./demodata/calendars.json")
	if err != nil {
		log.Fatal(err)
	}

	db := d.GetDB()
	err = db.Create(&calendars).Error
	if err != nil {
		return err
	}
	err = db.Create(&events).Error

	return
}

func parseDemodata(dest interface{}, path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		logError(err)
		return err
	}
	err = json.Unmarshal(bytes, &dest)
	if err != nil {
		logError(err)
	}
	return err
}

func (d *EventDemo) GetModel() *Event {
	sDate, _ := time.Parse("2006-01-02T15:04:05", d.StartDate)
	eDate, _ := time.Parse("2006-01-02T15:04:05", d.EndDate)

	event := Event{
		ID:          d.ID,
		EventUpdate: d.EventUpdate,
	}

	event.StartDate = &sDate
	event.EndDate = &eDate

	return &event
}
