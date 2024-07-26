package main

import (
  "log"
  "log/slog"
  "embed"
  "os"
  "net/http"

  "ftgodev-tut/handler"
  "ftgodev-tut/pkg/supabase"

  "github.com/go-chi/chi/v5"
  "github.com/joho/godotenv"
)

//go:embed public 
var FS embed.FS

func main() {
  if err := initEverything(); err != nil {
    log.Fatal(err)
  }

  router := chi.NewMux();
  router.Use(handler.WithAccess)

  router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
  router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
  router.Get("/login", handler.MakeHandler(handler.HandleLoginIndex))
  router.Post("/login", handler.MakeHandler(handler.HandleLogin))

  port := os.Getenv("PORT")
  slog.Info("server started", "port", port)
  log.Fatal(http.ListenAndServe(port, router))
}

// init everything
func initEverything() error {
  if err := godotenv.Load(); err != nil {
    return err
  }

  return supabase.Init()
}
