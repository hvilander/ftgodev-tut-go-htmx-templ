package handler

import (
  "fmt"
  "net/http"
  "encoding/json"
  "database/sql"
  "context"
  
  "ftgodev-tut/db"
  "ftgodev-tut/models"

	"github.com/uptrace/bun"
  "github.com/go-chi/chi/v5"
  "github.com/google/uuid"
  //"github.com/replicate/replicate-go"
)

const (
  succeeded = "succeeded"
  processing = "processing"
)

type ReplicateResp struct {
  Input struct {
    Prompt string `json:"prompt"`
  } `json:"input"`
  Status string `json:"status"`
  Output []string `json:"output"`
}

func HandleReplicateCallback(w http.ResponseWriter, r *http.Request) error {
  var resp ReplicateResp 
  if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
    return err
  }

  if resp.Status == processing {
    // we do nothing on the call back that just reports processing
    return nil
  }

 if resp.Status != succeeded {
   return fmt.Errorf("Replicate callback responded with a non ok status: %s", resp.Status)
 }

  /*
  2024/08/20 06:36:17
  ERROR internal server error err="Replicate callback invalid batch id invalid UUID length: 0" path=/replicate/callback/1803a731-16c0-4d1e-9a43-b8f4c4b59a86/6605ec8c-86a2-4b1d-bb75-a54c55110ae7
  */
 batchID, err := uuid.Parse(chi.URLParam(r, "batchID"))

 if err != nil {
   return fmt.Errorf("Replicate callback invalid batch id %s", err)
 }

 images, err := db.GetImagesByBatchID(batchID)
 if err != nil {
   return fmt.Errorf("Failed to find images with batchID: %s, err: %s", batchID, err)
 }

 if len(images) != len(resp.Output) {
   return fmt.Errorf("Replicate and db disagree about number of images")
 }

 err = db.Bun.RunInTx(
   r.Context(),
   &sql.TxOptions{},
   func(ctx context.Context, tx bun.Tx,) error {
      for i, imageURL := range resp.Output {
        images[i].Status = models.ImageStatusCompleted
        images[i].Location = imageURL 
        images[i].Prompt = resp.Input.Prompt 

        if err := db.UpdateImage(tx, &images[i]); err != nil {

        }
      }
      return nil
    },
  )

  fmt.Println(err)
  return err 
}
