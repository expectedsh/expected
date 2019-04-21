package apiserver

import (
	"fmt"
	"github.com/expectedsh/expected/pkg/apps/apiserver/request"
	"github.com/expectedsh/expected/pkg/apps/apiserver/response"
	"github.com/expectedsh/expected/pkg/models/images"
	"github.com/expectedsh/expected/pkg/protocol"
	"github.com/expectedsh/expected/pkg/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *App) ListImages(w http.ResponseWriter, r *http.Request) {
	account := request.GetAccount(r)
	imagesStats, err := images.FindImagesSummariesByNamespaceID(r.Context(), account.ID)
	if err != nil {
		logrus.WithError(err).WithField("account", account.ID).Error("unable to get images list")
		response.ErrorInternal(w)
		return
	}
	if imagesStats == nil {
		imagesStats = []*images.ImageSummary{}
	}
	response.Resource(w, "images", imagesStats)
}

func (s *App) GetImage(w http.ResponseWriter, r *http.Request) {
	account := request.GetAccount(r)
	name := mux.Vars(r)["name"]
	tag := mux.Vars(r)["tag"]
	image, err := images.FindImageDetail(r.Context(), account.ID, name, tag)
	if err != nil {
		logrus.WithError(err).WithField("account", account.ID).Error("unable find image detail")
		response.ErrorInternal(w)
		return
	}
	response.Resource(w, "image", image)
}

func (s *App) DeleteImage(w http.ResponseWriter, r *http.Request) {
	account := request.GetAccount(r)
	id := mux.Vars(r)["id"]
	if _, err := uuid.Parse(id); err != nil {
		response.ErrorBadRequest(w, "Invalid image id.", nil)
		return
	}
	img, err := images.FindImageByID(r.Context(), id)
	if err != nil {
		logrus.WithError(err).WithField("account", account.ID).Error("unable find image")
		response.ErrorInternal(w)
		return
	}
	if img == nil {
		response.ErrorNotFound(w)
		return
	}
	if img.NamespaceID != account.ID {
		response.ErrorForbidden(w)
		return
	}
	log := logrus.
		WithField("task", "api-delete-image").
		WithField("repo", fmt.Sprintf("%v/%v", img.NamespaceID, img.Name)).
		WithField("digest", img.Digest).
		WithField("tag", img.Tag).
		WithField("id", img.ID).
		WithField("tag", img.Tag)
	if img.DeleteMode {
		response.Error(w, http.StatusConflict, "The image is being deleted")
		return
	}
	if err := images.UpdateImageDeleteMode(r.Context(), img.ID); err != nil {
		log.WithError(err).Error("can't update image into delete mode")
		response.ErrorInternal(w)
		return
	}
	if _, err := services.Image().Client().DeleteImage(r.Context(), &protocol.DeleteImageRequest{
		Id: id,
	}); err != nil {
		log.WithError(err).Error("can't publish delete message")
		response.ErrorInternal(w)
		return
	}
	w.WriteHeader(202)
}
