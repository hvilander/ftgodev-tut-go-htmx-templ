package main

import (
  "log"
  "log/slog"
  "embed"
  "os"
  "net/http"

  "ftgodev-tut/handler"
  "ftgodev-tut/pkg/sb"
  "ftgodev-tut/db"

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
  router.Use(handler.WithUser) // user may exist 

  router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
  router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
  router.Get("/login", handler.MakeHandler(handler.HandleLoginIndex))
  router.Post("/login", handler.MakeHandler(handler.HandleLogin))
  router.Get("/login/provider/google", handler.MakeHandler(handler.HandleLoginWithGoogle))
  router.Post("/logout", handler.MakeHandler(handler.HandleLogout))
  router.Get("/signup", handler.MakeHandler(handler.HandleSignupIndex))
  router.Post("/signup", handler.MakeHandler(handler.HandleSignup))
  router.Get("/auth/callback", handler.MakeHandler(handler.HandleAuthCallback))

  // Auth required
  router.Group( func(auth chi.Router) {
    auth.Use(handler.WithAuth)
    auth.Get("/account/setup", handler.MakeHandler(handler.HandleAccountSetupIndex))
    auth.Post("/account/setup", handler.MakeHandler(handler.HandleAccountSetup))
  })

  // Auth and account required
  router.Group(func(auth chi.Router) {
    auth.Use(handler.WithAuth, handler.WithAccountSetup)
    auth.Get("/settings", handler.MakeHandler(handler.HandleSettingsIndex))
    auth.Put("/settings/account/profile", handler.MakeHandler(handler.HandleProfile))

    auth.Post("/auth/reset-password", handler.MakeHandler(handler.HandleResetPasswordCreate))
    auth.Put("/auth/reset-password", handler.MakeHandler(handler.HandleResetPasswordUpdate))
    auth.Get("/auth/reset-password", handler.MakeHandler(handler.HandleResetPassword))

    auth.Get("/generate", handler.MakeHandler(handler.HandleGenerateIndex))
  })

  port := os.Getenv("PORT")
  slog.Info("server started", "port", port)
  log.Fatal(http.ListenAndServe(port, router))
}

// init everything
func initEverything() error {
  if err := godotenv.Load(); err != nil {
    return err
  }

  if err := db.Init(); err != nil {
    return err
  }

  return sb.Init()
}
