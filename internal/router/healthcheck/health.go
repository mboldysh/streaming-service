package healthcheck

import (
	"net/http"

	"github.com/mboldysh/streaming-service/internal/router/writer"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writer.RespondWithJSON(w, http.StatusOK, nil)
}
