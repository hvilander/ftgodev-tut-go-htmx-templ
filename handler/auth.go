package handler

import (
  "net/http"
  "ftgodev-tut/view/auth"
  "ftgodev-tut/pkg/sb"
  "ftgodev-tut/pkg/util"
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

  if !util.IsValidEmail(creds.Email) {
    return render(r, w, auth.LoginForm(creds, auth.LoginErrors{
      Email: "malformed email addr",
    }))
  }

  if reason, ok := util.ValidatePassword(creds.Password); !ok {
    return render(r, w, auth.LoginForm(creds, auth.LoginErrors{
      Password: reason,
    }))
  }


   resp, err := sb.Client.Auth.SignIn(r.Context(), creds)

   if err != nil {
     return render(r, w, auth.LoginForm(creds, auth.LoginErrors{
       InvalidCreds: "the creds are bad",
     }))
   }

   // set a cookie
   cookie := &http.Cookie{
     Value: resp.AccessToken,
     Name: "at",
     HttpOnly: true,
     Secure: true,
   }

   http.SetCookie(w, cookie);
   http.Redirect(w, r, "/", http.StatusSeeOther)

  return nil
}
