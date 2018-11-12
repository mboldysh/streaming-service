package userrouter

import (
	"github.com/go-chi/chi"
	"github.com/mboldysh/streaming-service/internal/service"
	"github.com/mboldysh/streaming-service/pkg/router"
)

type trackHandler struct {
	trackService service.TrackService
}

func New(trackService service.TrackService) router.Router {
	h := &trackHandler{
		trackService: trackService,
	}
	return h.initRoutes()
}

func (s *trackHandler) initRoutes() router.Router {
	r := chi.NewRouter()
	r.Post("/{userID}/tracks", s.Upload)
	r.Get("/{userID}/tracks", s.FindAll)
	r.Get("/{userID}/tracks/{trackName}", s.GetPresignedURL)
	return router.New("/users", r)
}
