package data

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"web-widgets/scheduler-go/common"
)

type EventDemo struct {
	ID        common.FuzzyInt `json:"id"`
	Name      string          `json:"text"`
	StartDate string          `json:"start_date"`
	EndDate   string          `json:"end_date"`
	AllDay    bool            `json:"allDay"`
	Type      string          `json:"type"`
	Details   string          `json:"details"`
}

func (d *EventDemo) GetModel() *Event {
	sDate, _ := time.Parse("2006-01-02 15:04:05", d.StartDate)
	eDate, _ := time.Parse("2006-01-02 15:04:05", d.EndDate)

	return &Event{
		ID:        int(d.ID),
		Name:      d.Name,
		StartDate: &sDate,
		EndDate:   &eDate,
		AllDay:    d.AllDay,
		Type:      d.Type,
		Details:   d.Details,
	}
}

func getData() ([]EventDemo, error) {
	bytes, err := ioutil.ReadFile("./demodata/events.json")
	if err != nil {
		logError(err)
		return nil, err
	}
	data := make([]EventDemo, 0)
	err = json.Unmarshal(bytes, &data)

	return data, err
}

func dataDown(d *DAO) {
	d.mustExec("DELETE from events")
}

func dataUp(d *DAO) error {
	db := d.GetDB()
	events, err := getData()
	if err != nil {
		logError(err)
		return err
	}

	tx := db.Begin()
	for _, r := range events {
		err = tx.Create(r.GetModel()).Error
		if err != nil {
			return err
		}
	}
	return tx.Commit().Error
}
