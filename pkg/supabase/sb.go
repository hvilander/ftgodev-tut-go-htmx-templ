package supabase

import (
  "os"

  "github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func Init() error {
  sbHost := os.Getenv("SB_URL")
  sbSecret := os.Getenv("SB_SECRET")
  Client = supabase.CreateClient(sbHost, sbSecret)
  return nil
}
