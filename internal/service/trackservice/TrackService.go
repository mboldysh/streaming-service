package trackservice

import (
	"fmt"
	"strings"

	"github.com/mboldysh/streaming-service/internal/model"
	"github.com/mboldysh/streaming-service/internal/service"
	"github.com/mboldysh/streaming-service/internal/store"
)

type trackService struct {
	trackStore store.TrackStore
}

//New initialize a new TrackService
func New(trackStore store.TrackStore) service.TrackService {
	return &trackService{
		trackStore: trackStore,
	}
}

//Upload uploads track to bucket
func (s trackService) Upload(track model.UploadTrack, userID string) error {
	key := genKey(userID, track.Name)
	return s.trackStore.Upload(key, track.File)
}

//FindAll finds all tracks by user id
func (s trackService) FindAll(userID string) ([]model.Track, error) {
	tracks, err := s.trackStore.FindAll(userID)

	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("%s/", userID)

	for index, track := range tracks {
		tracks[index].Name = strings.TrimPrefix(track.Name, prefix)
	}
	return tracks, err
}

//GetPresignedURL creates presigned URL
func (s trackService) GetPresignedURL(userID, trackName string) (*model.PresignedTrack, error) {
	key := genKey(userID, trackName)
	return s.trackStore.GetPresignedURL(key, trackName)
}

//genKey generates bucket key
func genKey(userID, trackName string) string {
	return fmt.Sprintf("%s/%s", userID, trackName)
}
