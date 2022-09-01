package data

import (
	"time"
	"web-widgets/scheduler-go/common"

	"gorm.io/gorm"
)

type EventUpdate struct {
	Name         string          `json:"text"`
	StartDate    *time.Time      `json:"start_date"`
	EndDate      *time.Time      `json:"end_date"`
	AllDay       bool            `json:"allDay"`
	Type         common.FuzzyInt `json:"type"`
	Details      string          `json:"details"`
	AttachedData []*BinaryData   `json:"files"`
}

type EventsDAO struct {
	db *gorm.DB
}

func NewEventsDAO(db *gorm.DB) *EventsDAO {
	return &EventsDAO{db}
}

func (d *EventsDAO) GetOne(id int) (Event, error) {
	event := Event{}
	err := d.db.Preload("AttachedData", func(db *gorm.DB) *gorm.DB {
		return d.db.Order("binary_data.id ASC")
	}).Find(&event, id).Error

	return event, err
}

func (d *EventsDAO) GetAll() ([]Event, error) {
	events := make([]Event, 0)
	err := d.db.
		Preload("AttachedData", func(db *gorm.DB) *gorm.DB {
			return d.db.Order("binary_data.id ASC")
		}).
		Find(&events).Error

	return events, err
}

func (d *EventsDAO) Add(update *EventUpdate) (int, error) {
	event := Event{}
	update.fillModel(&event)

	err := d.db.Create(&event).Error
	return event.ID, err
}

func (d *EventsDAO) Update(id int, update *EventUpdate) error {
	event := Event{}
	err := d.db.Find(&event, id).Error
	if err != nil {
		return err
	}

	update.fillModel(&event)
	err = d.db.Save(&event).Error
	if err == nil {
		err = d.db.Model(&BinaryData{}).Where("event_id = ?", event.ID).Update("event_id", nil).Error
		if err != nil {
			return err
		}

		if len(update.AttachedData) > 0 {
			tempIDs := make([]int, len(update.AttachedData))
			for i, x := range update.AttachedData {
				tempIDs[i] = x.ID
			}
			err = d.db.Model(&BinaryData{}).Where("id in (?)", tempIDs).Update("event_id", event.ID).Error
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (d *EventsDAO) Delete(id int) error {
	err := d.db.Where("event_id = ?", id).Delete(&BinaryData{}).Error
	if err == nil {
		err = d.db.Delete(&Event{}, id).Error
	}
	return err
}

func (d *EventUpdate) fillModel(ev *Event) {
	if ev != nil {
		ev.Name = d.Name
		ev.StartDate = d.StartDate
		ev.EndDate = d.EndDate
		ev.AllDay = d.AllDay
		ev.Type = d.Type
		ev.Details = d.Details
	}
}
