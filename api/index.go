package api

import (
	"context"
	"errors"
	"net/http"

	remote "github.com/mkozhukh/go-remote"

	"web-widgets/scheduler-go/data"
)

type UserID int
type DeviceID int

type EventConfig struct {
	From  int        `json:"-"`
	Event data.Event `json:"event"`
	Type  string     `json:"type"`
}

func BuildAPI(db *data.DAO) *remote.Server {
	if remote.MaxSocketMessageSize < 32000 {
		remote.MaxSocketMessageSize = 32000
	}

	api := remote.NewServer(&remote.ServerConfig{
		WebSocket: true,
	})

	api.Events.AddGuard("events", func(m *remote.Message, c *remote.Client) bool {
		tm, ok := m.Content.(EventConfig)
		if !ok {
			return false
		}

		return int(tm.From) != c.ConnID
	})

	api.Connect = func(r *http.Request) (context.Context, error) {
		id, _ := r.Context().Value("user_id").(int)
		if id == 0 {
			return nil, errors.New("access denied")
		}
		device, _ := r.Context().Value("device_id").(int)
		if device == 0 {
			return nil, errors.New("access denied")
		}

		return context.WithValue(
			context.WithValue(r.Context(), remote.UserValue, id),
			remote.ConnectionValue, device), nil
	}

	return api
}
