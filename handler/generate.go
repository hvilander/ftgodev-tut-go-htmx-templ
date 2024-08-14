package handler

import (
  "net/http"
  "ftgodev-tut/view/generate"
  "ftgodev-tut/models"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
  images := make([]models.Image, 20)
  data := generate.ViewData{Images: images}

  images[0].Status = models.ImageStatusPending
  return render(r, w, generate.Index(data))
}
