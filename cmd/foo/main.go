package main
/* playground to experiment with */

import (
  "context"
  "fmt"
  "log"

  "github.com/replicate/replicate-go"
  "github.com/joho/godotenv"
)


func main() {
  if err := godotenv.Load(); err != nil {
    log.Fatal(err)
  }
  ctx := context.Background()

  r8, err := replicate.NewClient(replicate.WithTokenFromEnv())
  if err != nil {
    log.Fatal("error setting up replicate client", err)
  }

  version := "ac732df83cea7fff18b8472768c88ad041fa750ff7682a21affe81863cbe77e4"


  input := replicate.PredictionInput {
    "prompt": "sexy lambo daddy fighting a rancor",
  }

  // from webhook.site
  webhook := replicate.Webhook{
    URL: "https://webhook.site/fff3ee71-9efe-41b4-8f66-93dda162e2af",
    Events: []replicate.WebhookEventType{"start", "completed"},
  }

  //run a model and wait for its output
  output, err := r8.CreatePrediction(ctx, version, input, &webhook, false)
  if err != nil {
    log.Fatal("error runing model from replicate", err)
  }
  fmt.Println("output: ", output)



}
