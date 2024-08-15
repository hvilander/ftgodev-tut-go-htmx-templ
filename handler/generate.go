package handler

import (
  "strconv"
  "net/http"
  "ftgodev-tut/view/generate"
  "ftgodev-tut/models"
  "ftgodev-tut/db"
  "github.com/go-chi/chi/v5"
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
  id, err := strconv.Atoi(chi.URLParam(r, "id"))
  if err != nil {
    return err
  }

  image, err := db.GetImageByID(id)
  if err != nil {
    return err
  }


  return render(r, w, generate.GalleryImage(image))
}

func HandleGenerateCreate(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r)
  prompt := "sexy lambo"
  img := models.Image{
    Prompt: prompt,
    UserID: user.ID,
    Status: 1,
  }

  if err := db.CreateImage(&img); err != nil {
    return err
  }

  return render(r, w, generate.GalleryImage(img))
}

