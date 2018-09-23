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

func New(trackStore store.TrackStore) service.TrackService {
	return &trackService{
		trackStore: trackStore,
	}
}

func (service trackService) Upload(track model.UploadTrack, userID string) error {
	key := genKey(userID, track.Name)
	return service.trackStore.Upload(key, track.File)
}

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

func (s trackService) GetPresignedURL(userID, trackName string) (*model.PresignedTrack, error) {
	key := genKey(userID, trackName)
	return s.trackStore.GetPresignedURL(key, trackName)
}

func genKey(userID, trackName string) string {
	return fmt.Sprintf("%s/%s", userID, trackName)
}
