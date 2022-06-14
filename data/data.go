package data

type Event struct {
	ID int `json:"id"`
	EventUpdate
}

type Calendar struct {
	ID int `json:"id"`
	CalendarUpdate
}
