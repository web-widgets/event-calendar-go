package main

import "web-widgets/scheduler-go/data"

type ConfigServer struct {
	URL  string
	Port string
	Cors []string
}

type AppConfig struct {
	Server ConfigServer
	DB     data.DBConfig
}
