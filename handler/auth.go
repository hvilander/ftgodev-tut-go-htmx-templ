package handler

import (
  "net/http"
  "ftgodev-tut/view/auth"
  "github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
  return auth.Login().Render(r.Context(), w)
}


func HandleLogin(w http.ResponseWriter, r *http.Request) error {
  creds := supabase.UserCredentials{
    Email: r.FormValue("email"),
    Password: r.FormValue("password"),
  }

  // call supabase



  return render(r, w, auth.LoginForm(creds, auth.LoginErrors{
    InvalidCreds: "the creds are bad",
  }))
}
