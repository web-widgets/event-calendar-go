package data

import (
	"gorm.io/gorm"
)

type CalendarUpdate struct {
	Name        string `json:"label"`
	Active      bool   `json:"active"`
	Color       *Color `json:"color"`
	Description string `json:"description"`
}

type CalendarsDAO struct {
	db *gorm.DB
}

func NewCalendarsDAO(db *gorm.DB) *CalendarsDAO {
	return &CalendarsDAO{db}
}

func (d *CalendarsDAO) GetOne(id int) (Calendar, error) {
	calendar := Calendar{}
	err := d.db.Find(&calendar, id).Error

	return calendar, err
}

func (d *CalendarsDAO) GetAll() ([]Calendar, error) {
	calendars := make([]Calendar, 0)
	err := d.db.Find(&calendars).Error

	return calendars, err
}

func (d *CalendarsDAO) Add(update *CalendarUpdate) (int, error) {
	calendar := Calendar{}
	update.fillModel(&calendar)
	calendar.Active = true

	err := d.db.Create(&calendar).Error
	return calendar.ID, err
}

func (d *CalendarsDAO) Update(id int, update *CalendarUpdate) error {
	calendar := Calendar{}
	err := d.db.Find(&calendar, id).Error
	if err != nil {
		return err
	}

	update.fillModel(&calendar)
	err = d.db.Save(&calendar).Error
	return err
}

func (d *CalendarsDAO) Delete(id int) error {
	err := d.db.Delete(&Calendar{}, id).Error
	if err == nil {
		err = d.db.Where("type = ?", id).Delete(&Event{}).Error
	}
	return err
}

func (d *CalendarUpdate) fillModel(ev *Calendar) {
	if ev != nil {
		ev.Name = d.Name
		ev.Active = d.Active
		ev.Color = d.Color
		ev.Description = d.Description
	}
}
