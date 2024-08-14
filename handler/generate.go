package handler

import (
  "net/http"
  "ftgodev-tut/view/generate"
  "ftgodev-tut/models"
  "github.com/go-chi/chi/v5"
  "log/slog"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
  //images := make([]models.Image, 20)
  data := generate.ViewData{Images: []models.Image{}}

  //images[0].Status = models.ImageStatusPending
  return render(r, w, generate.Index(data))
}

func HandleGenerateImageStatus(w http.ResponseWriter, r *http.Request) error {
  id := chi.URLParam(r, "id")
  slog.Info("EHLO from image status", "id", id)

  // todo fetch from db
  image := models.Image{
    Status: models.ImageStatusPending,
  }
  return render(r, w, generate.GalleryImage(image))
}

func HandleGenerateCreate(w http.ResponseWriter, r *http.Request) error {
  return render(r, w, generate.GalleryImage(models.Image{Status: models.ImageStatusPending}))
}
