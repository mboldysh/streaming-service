package service

import (
	"github.com/mboldysh/streaming-service/internal/model"
)

type TrackService interface {
	Upload(track model.UploadTrack, userID string) error
	FindAll(userID string) ([]model.Track, error)
	GetPresignedURL(userID, trackName string) (*model.PresignedTrack, error)
}
