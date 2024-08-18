package handler


import (
	"github.com/uptrace/bun"
  "context"
  "database/sql"
  "strconv"
  "net/http"
  "ftgodev-tut/view/generate"
  "ftgodev-tut/models"
  "ftgodev-tut/db"
  "github.com/go-chi/chi/v5"
  "github.com/google/uuid"
  "ftgodev-tut/pkg/kit/validate"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r) 
  images, err := db.GetImagesByUserID(user.ID)
  if err != nil {
    return err
  }

  data := generate.ViewData{
    Images: images,
    FormParams: generate.FormParams{},
    FormErrors: generate.FormErrors{},
  }
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
  amount, _ := strconv.Atoi(r.FormValue("amount"))

  params := generate.FormParams{
    Prompt: r.FormValue("prompt"),
    Amount: amount,
  }

  var errors generate.FormErrors

  if amount <= 0 || amount > 8 {
    errors.Amount = "invalid amount"
  }

  ok := validate.New(params, validate.Fields{
    "Prompt": validate.Rules(validate.Min(10), validate.Max(100)),
  }).Validate(&errors)

  if !ok {
    return render(r,w, generate.Form(params, errors))
  }

  err := db.Bun.RunInTx(
    r.Context(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
      batchID := uuid.New()
      for i :=0; i < amount; i++ {
        img := models.Image{
          Prompt: params.Prompt,
          UserID: user.ID,
          Status: models.ImageStatusPending,
          BatchID: batchID,
        }
        if err := db.CreateImage(&img); err != nil {
          return err
        }
      }
      return nil
  })

  if err != nil {
    return nil
  }

  return hxRedirect(w, r, "/generate")
}

