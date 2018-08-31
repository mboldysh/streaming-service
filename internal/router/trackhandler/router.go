package trackhandler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mboldysh/streaming-service/internal/service"
)

type TrackHandler struct {
	trackService service.TrackService
}

func New(trackService service.TrackService) http.Handler {
	h := &TrackHandler{
		trackService: trackService,
	}
	return h.initRoutes()
}

func (s *TrackHandler) initRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/users/{userID}/tracks", s.Upload)
	r.Get("/users/{userID}/tracks", s.FindAll)
	r.Get("/users/{userID}/tracks/{trackName}", s.GetPresignedURL)
	return r
}
