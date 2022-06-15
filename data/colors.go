package data

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Color struct {
	Background string `json:"background"`
	Border     string `json:"border"`
}

func (c *Color) Scan(value interface{}) error {
	if value == nil {
		*c = Color{}
		return nil
	}

	if s, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := s.(string); ok {
			err = json.Unmarshal([]byte(v), c)
			return nil
		}
	}

	return errors.New("failed to scan Color")
}

func (c *Color) Value() (driver.Value, error) {
	data, err := json.Marshal(c)
	return string(data), err
}
