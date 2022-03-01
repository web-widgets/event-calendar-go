package api

import (
	remote "github.com/mkozhukh/go-remote"

	"web-widgets/scheduler-go/data"
)

func BuildAPI(db *data.DAO) *remote.Server {
	if remote.MaxSocketMessageSize < 32000 {
		remote.MaxSocketMessageSize = 32000
	}

	api := remote.NewServer(&remote.ServerConfig{
		WebSocket: true,
	})

	return api
}
