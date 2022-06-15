package api

import "web-widgets/scheduler-go/data"

type EventType interface {
	GetFrom() int
}

type EventItem struct {
	From  int         `json:"-"`
	Type  string      `json:"type"`
	Event *data.Event `json:"event"`
}

type EventCalendar struct {
	From     int            `json:"-"`
	Type     string         `json:"type"`
	Calendar *data.Calendar `json:"calendar"`
}

func (e EventItem) GetFrom() int {
	return e.From
}

func (e EventCalendar) GetFrom() int {
	return e.From
}
