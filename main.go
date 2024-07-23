package main

import (
  "log"
  "log/slog"
  "os"
  "net/http"

  "ftgodev-tut/handler"

  "github.com/go-chi/chi/v5"
  "github.com/joho/godotenv"
)

func main() {
  if err := initEverything(); err != nil {
    log.Fatal(err)
  }

  router := chi.NewMux();

  router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

  port := os.Getenv("PORT")
  slog.Info("server started", "port", port)
  log.Fatal(http.ListenAndServe(port, router))
}


// init everything
func initEverything() error {
  return godotenv.Load();

  /*
  if err := godotenv.Load(); err != nil {
    return err
  }

  return nil
  */
}
