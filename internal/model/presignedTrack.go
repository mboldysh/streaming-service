package model

import (
	"net/http"
)

//PresignedTrack type details
type PresignedTrack struct {
	Name   string      `json:"name"`
	URL    string      `json:"url"`
	Header http.Header `json:"header"`
}
