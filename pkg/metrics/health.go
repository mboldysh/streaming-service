package metrics

import (
	"net/http"

	"github.com/mboldysh/streaming-service/pkg/httpwriter"
	"github.com/mboldysh/streaming-service/pkg/router"
)

type healthcheck struct{}

func NewHealthCheck() router.Router {
	return healthcheck{}.initRoutes()
}

func (h healthcheck) initRoutes() router.Router {
	return router.NewRouter("/metrics/health", h.healthHandler)
}

func (h healthcheck) healthHandler(w http.ResponseWriter, r *http.Request) {
	httpwriter.RespondWithJSON(w, http.StatusOK, nil)
}
