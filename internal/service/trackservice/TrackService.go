package trackservice

import (
	"fmt"

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
	return s.trackStore.FindAll(userID)
}

func (s trackService) GetPresignedURL(userID, trackName string) (*model.PresignedTrack, error) {
	key := genKey(userID, trackName)
	return s.trackStore.GetPresignedURL(key, trackName)
}

func genKey(userID, trackName string) string {
	return fmt.Sprintf("%s/%s", userID, trackName)
}
