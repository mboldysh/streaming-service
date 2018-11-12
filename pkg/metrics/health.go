package metrics

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mboldysh/streaming-service/pkg/httpwriter"
	"github.com/mboldysh/streaming-service/pkg/router"
)

type healthcheck struct{}

func NewHealthCheck() router.Router {
	return healthcheck{}.initRoutes()
}

func (h healthcheck) initRoutes() router.Router {
	r := chi.NewRouter()
	r.Get("/health", h.healthHandler)
	return router.New("/metrics", r)
}

func (h healthcheck) healthHandler(w http.ResponseWriter, r *http.Request) {
	httpwriter.RespondWithJSON(w, http.StatusOK, nil)
}
