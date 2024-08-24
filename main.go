package main

import (
  "log"
  "log/slog"
  "os"
  "net/http"

  "ftgodev-tut/handler"
  "ftgodev-tut/pkg/sb"
  "ftgodev-tut/db"

  "github.com/go-chi/chi/v5"
  "github.com/joho/godotenv"
)


func main() {
  if err := initEverything(); err != nil {
    log.Fatal(err)
  }

  router := chi.NewMux();
  router.Use(handler.WithUser) // user may exist 

  router.Handle("/*", public())
  router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
  router.Get("/login", handler.MakeHandler(handler.HandleLoginIndex))
  router.Post("/login", handler.MakeHandler(handler.HandleLogin))
  router.Get("/login/provider/google", handler.MakeHandler(handler.HandleLoginWithGoogle))
  router.Post("/logout", handler.MakeHandler(handler.HandleLogout))
  router.Get("/signup", handler.MakeHandler(handler.HandleSignupIndex))
  router.Get("/auth/callback", handler.MakeHandler(handler.HandleAuthCallback))
  router.Post("/replicate/callback/{userID}/{batchID}", handler.MakeHandler(handler.HandleReplicateCallback))

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
    auth.Get("/generate", handler.MakeHandler(handler.HandleGenerateIndex))
    auth.Post("/generate", handler.MakeHandler(handler.HandleGenerateCreate))
    auth.Get("/generate/image/status/{id}", handler.MakeHandler(handler.HandleGenerateImageStatus))
    auth.Get("/buy-credits", handler.MakeHandler(handler.HandleCreditsIndex))
    auth.Get("/checkout/create/{priceID}", handler.MakeHandler(handler.HandleStripeCheckoutCreate))
    auth.Get("/checkout/success/{sessionID}", handler.MakeHandler(handler.HandleStripeCheckoutSuccess))
    auth.Get("/checkout/cancel", handler.MakeHandler(handler.HandleStripeCheckoutCancel))

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
