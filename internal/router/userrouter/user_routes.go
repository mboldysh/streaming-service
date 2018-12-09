package userrouter

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mboldysh/streaming-service/internal/model"
	"github.com/mboldysh/streaming-service/pkg/httpwriter"
)

func (h *userRouter) upload(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	reader, err := r.MultipartReader()

	if err != nil {
		httpwriter.RespondWithError(w, http.StatusBadRequest, "Can't upload files")
		return
	}

	for {
		part, err := reader.NextPart()

		if err != nil {
			if err == io.EOF {
				break
			}
			httpwriter.RespondWithError(w, http.StatusBadRequest, "Can't upload files")
			return
		}

		if part.FileName() == "" {
			continue
		}

		track := model.UploadTrack{
			File: part,
			Name: part.FileName(),
		}

		err = h.trackService.Upload(track, userID)

		if err != nil {
			httpwriter.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	httpwriter.RespondWithJSON(w, http.StatusOK, nil)
}

func (h *userRouter) findAll(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	listTracks, err := h.trackService.FindAll(userID)
	if err != nil {
		httpwriter.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpwriter.RespondWithJSON(w, http.StatusOK, listTracks)
}

func (h *userRouter) getPresignedURL(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	trackName := chi.URLParam(r, "trackName")
	presignedTrack, err := h.trackService.GetPresignedURL(userID, trackName)

	if err != nil {
		httpwriter.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	httpwriter.RespondWithJSON(w, http.StatusOK, presignedTrack)
}

func (h *userRouter) deleteObject(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	trackName := chi.URLParam(r, "trackName")

	listTracks, err := h.trackService.DeleteObject(userID, trackName)

	if err != nil {
		httpwriter.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpwriter.RespondWithJSON(w, http.StatusOK, listTracks)
}
