package main

import (
	"time"
)

type Stop struct {
	Id        string `json:"id"`
	Indicator string `json:"indicator"`
	Name      string `json:"name"`
	Lines     []Line `json:"lines"`
}

type Line struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Vehicle struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Destination string `json:"destination"`
}

type Arrival struct {
	Vehicle  Vehicle   `json:"vehicle"`
	Expected time.Time `json:"expected"`
}

type StopsService interface {
	Get(lat float64, lon float64, radius uint) ([]Stop, error)
}

type ArrivalsService interface {
	Get(stopId string) ([]Arrival, error)
}
