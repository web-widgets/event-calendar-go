package data

type Event struct {
	ID int `json:"id"`
	EventUpdate
}

type Calendar struct {
	ID int `json:"id"`
	CalendarUpdate
}

type BinaryData struct {
	ID   int    `json:"id"`
	Path string `json:"-"`
	Name string `json:"name"`
	URL  string `json:"url"`

	EventID int    `json:"-"`
	Event   *Event `json:"-"`
}
