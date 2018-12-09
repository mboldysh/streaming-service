package metrics

import (
	"net/http"

	"github.com/mboldysh/streaming-service/pkg/httpwriter"
	"github.com/mboldysh/streaming-service/pkg/router"
)

type healthcheck struct{}

// func NewHealthCheck() router.Router {
// 	return healthcheck{}.initRoutes()
// }

// func (h healthcheck) initRoutes() router.Endpoint {
// 	return router.Get("/metrics/health", h.healthHandler)
// }

func HealthCheck() router.Endpoint {
	return router.Get("/metrics/health", healthHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	httpwriter.RespondWithJSON(w, http.StatusOK, nil)
}
