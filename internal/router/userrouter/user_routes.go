package userrouter

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mboldysh/streaming-service/internal/model"
	"github.com/mboldysh/streaming-service/pkg/httpwriter"
)

func (h *userRouter) upload(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	file, handler, err := r.FormFile("uploadFile")

	if err != nil {
		httpwriter.RespondWithError(w, http.StatusBadRequest, "Can't upload file")
		return
	}
	defer file.Close()

	track := model.UploadTrack{
		File: file,
		Name: handler.Filename,
	}

	err = h.trackService.Upload(track, userID)
	if err != nil {
		httpwriter.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	httpwriter.RespondWithJSON(w, http.StatusOK, nil)
}

func (h *userRouter) findAll(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	listTracks, err := h.trackService.FindAll(userID)
	if err != nil {
		httpwriter.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	httpwriter.RespondWithJSON(w, http.StatusOK, listTracks)
}

func (h *userRouter) getPresignedURL(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	trackName := chi.URLParam(r, "trackName")
	presignedTrack, err := h.trackService.GetPresignedURL(userID, trackName)

	if err != nil {
		httpwriter.RespondWithError(w, http.StatusNotFound, err.Error())
	}
	httpwriter.RespondWithJSON(w, http.StatusOK, presignedTrack)
}
