package handler


import (
  "fmt"
  "context"
  "database/sql"
  "strconv"
  "net/http"
  "log"

  "ftgodev-tut/view/generate"
  "ftgodev-tut/models"
  "ftgodev-tut/db"
  "ftgodev-tut/pkg/kit/validate"

  "github.com/uptrace/bun"
  "github.com/replicate/replicate-go"
  "github.com/go-chi/chi/v5"
  "github.com/google/uuid"
)


const creditsPerImage = 2

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

  creditsNeeded := params.Amount * creditsPerImage
  if user.Account.Credits < creditsNeeded {
    errors.AdditionalCreditsNeeded = creditsNeeded - user.Account.Credits
    return render(r,w, generate.Form(params, errors))
  }

  user.Account.Credits -= creditsNeeded
  if err := db.UpdateAccount(&user.Account); err != nil {
    return err
  }

  batchID := uuid.New()
  genParams := GenerateImageParams{
    Prompt: params.Prompt,
    Amount: params.Amount,
    UserID: user.ID,
    BatchID: batchID,
  }

  genErr:= generateImages(r.Context(), genParams)
  if genErr != nil {
    return genErr 
  }

  err := db.Bun.RunInTx( r.Context(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
    for i :=0; i < amount; i++ {
        img := models.Image{
          UserID: user.ID,
          Status: models.ImageStatusPending,
          BatchID: batchID,
        }
        if err := db.CreateImage(tx, &img); err != nil {
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

type GenerateImageParams struct {
  Prompt string
  Amount int
  BatchID uuid.UUID
  UserID uuid.UUID
}

func generateImages(ctx context.Context, params GenerateImageParams) error {
  r8, err := replicate.NewClient(replicate.WithTokenFromEnv())
  if err != nil {
    log.Fatal("error setting up replicate client", err)
  }

  version := "ac732df83cea7fff18b8472768c88ad041fa750ff7682a21affe81863cbe77e4"

  input := replicate.PredictionInput {
    "prompt": params.Prompt,
    "num_outputs": params.Amount,
  }

  baseURL := "https://webhook.site/fff3ee71-9efe-41b4-8f66-93dda162e2af" 
  url := fmt.Sprintf("%s/%s/%s", baseURL, params.UserID, params.BatchID)

  // from webhook.site
  webhook := replicate.Webhook{
    URL: url,
    Events: []replicate.WebhookEventType{"completed"},
  }

  //run a model and wait for its output
  _, err = r8.CreatePrediction(ctx, version, input, &webhook, false)

  if err != nil {
    return err
  }
  return nil

}
