package data

import (
	"time"

	"gorm.io/gorm"
)

func NewEventsDAO(db *gorm.DB) *EventsDAO {
	return &EventsDAO{db}
}

type EventUpdate struct {
	Name      string     `json:"text"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	AllDay    bool       `json:"allDay"`
	Type      string     `json:"type"`
	Details   string     `json:"details"`
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

func (d *EventsDAO) Add(update *EventUpdate) (int, error) {
	event := Event{}
	update.FillModel(&event)

	err := d.db.Save(&event).Error
	return event.ID, err
}

func (d *EventsDAO) Update(id int, update *EventUpdate) error {
	event := Event{}
	err := d.db.Find(&event, id).Error
	if err != nil || event.ID == 0 {
		return err
	}

	update.FillModel(&event)
	err = d.db.Save(&event).Error
	return err
}

func (d *EventsDAO) Delete(id int) error {
	err := d.db.Delete(&Event{}, id).Error
	return err
}

func (d *EventUpdate) FillModel(ev *Event) {
	if ev != nil {
		ev.Name = d.Name
		ev.StartDate = d.StartDate
		ev.EndDate = d.EndDate
		ev.AllDay = d.AllDay
		ev.Type = d.Type
		ev.Details = d.Details
	}
}
