package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-widgets/scheduler-go/data"

	"github.com/go-chi/chi"
)

type Response struct {
	ID int `json:"id"`
}

func NumberParam(r *http.Request, key string) int {
	value := chi.URLParam(r, key)
	num, _ := strconv.Atoi(value)

	return num
}

func ParseFormEvent(w http.ResponseWriter, r *http.Request) (data.EventUpdate, error) {
	c := data.EventUpdate{}

	body := http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(body)
	err := dec.Decode(&c)

	return c, err
}
