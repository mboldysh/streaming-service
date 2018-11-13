package model

import (
	"io"
)

//UploadTrack model
type UploadTrack struct {
	Name string
	File io.Reader
}
