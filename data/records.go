package data

import (
	"time"
	"web-widgets/scheduler-go/common"

	"gorm.io/gorm"
)

func NewEventsDAO(db *gorm.DB) *EventsDAO {
	return &EventsDAO{db}
}

type EventUpdate struct {
	ID        common.FuzzyInt `json:"id"`
	Name      string          `json:"text"`
	StartDate *time.Time      `json:"start_date"`
	EndDate   *time.Time      `json:"end_date"`
	Readonly  bool            `json:"readonly"`
	AllDay    bool            `json:"allDay"`
	Type      string          `json:"type"`
	Details   string          `json:"details"`
}

type EventsDAO struct {
	db *gorm.DB
}


func (d *EventsDAO) GetOne(id int) (Event, error) {
	event := Event{}
	err := d.db.Find(&event, id).Error

	return event, err
}

func (d *EventsDAO) GetAll() ([]Event, error) {
	events := make([]Event, 0)
	err := d.db.Find(&events).Error

	return events, err
}

func (d *EventsDAO) Add(event *Event) (int, error) {
	event.ID = 0
	err := d.db.Save(event).Error
	return event.ID, err
}

func (d *EventsDAO) Update(event *Event) error {
	c := Event{}
	err := d.db.Find(&c, event.ID).Error
	if err != nil || c.ID == 0 {
		return err
	}
	err = d.db.Save(&event).Error
	return err
}

func (d *EventsDAO) Delete(id int) error {
	err := d.db.Delete(&Event{}, id).Error
	return err
}

func (d *EventUpdate) GetModel() *Event {
	return &Event{
		ID:        int(d.ID),
		Name:      d.Name,
		StartDate: d.StartDate,
		EndDate:   d.EndDate,
		Readonly:  d.Readonly,
		Type:      d.Type,
		Details:   d.Details,
	}
}
