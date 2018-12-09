package userrouter

import (
	"github.com/go-chi/chi"
	"github.com/mboldysh/streaming-service/internal/service"
	"github.com/mboldysh/streaming-service/pkg/router"
)

//userRouter is a router to talk with the user controller
type userRouter struct {
	trackService service.TrackService
}

//New initialize a new user router
func New(trackService service.TrackService) router.Router {
	h := &userRouter{
		trackService: trackService,
	}
	return h.initRoutes()
}

//initRoutes initializes routes in user router
func (s *userRouter) initRoutes() router.Router {
	r := chi.NewRouter()
	r.Post("/{userID}/tracks", s.upload)
	r.Get("/{userID}/tracks", s.findAll)
	r.Get("/{userID}/tracks/{trackName}", s.getPresignedURL)
	return router.NewRouter("/users", r)
}
