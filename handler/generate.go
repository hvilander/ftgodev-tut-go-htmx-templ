package handler

import (
  "net/http"
  "ftgodev-tut/view/generate"
  "ftgodev-tut/models"
  "ftgodev-tut/db"
  "github.com/go-chi/chi/v5"
  "log/slog"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r) 
  images, err := db.GetImagesByUserID(user.ID)
  if err != nil {
    return err
  }

  data := generate.ViewData{Images: images}
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

func CreateImage(image *models.Image) err {
  _, err := Bun.NewInsert().
    Model(image).
    Exec(context.Background())
  return err
}

func HandleGenerateCreate(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r)
  prompt := "sexy lambo"
  img := models.Image{
    Prompt: prompt,
    UserID: user.ID,
  }

  return render(r, w, generate.GalleryImage(models.Image{Status: models.ImageStatusPending}))
}
