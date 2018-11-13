package httpwriter

import (
	"encoding/json"
	"net/http"
)

//RespondWithJSON send HTTP response with specified payload in JSON format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//RespondWithError send HTTP response with specifiend HTTP code and payload
//in JSON format
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"message": msg})
}
