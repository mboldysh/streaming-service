package store

import (
	"io"

	"github.com/mboldysh/streaming-service/internal/model"
)

//TrackStore representation
type TrackStore interface {
	Upload(key string, file io.Reader) error
	FindAll(userID string) ([]model.Track, error)
	GetPresignedURL(key, trackName string) (*model.PresignedTrack, error)
	DeleteObject(key string) error
}
