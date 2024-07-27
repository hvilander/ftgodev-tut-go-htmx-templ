package handler

import (
  "fmt"
  "log/slog"
  "net/http"
  "ftgodev-tut/view/auth"
  "ftgodev-tut/pkg/sb"
  "ftgodev-tut/pkg/kit/validate"
  "github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
  return auth.Login().Render(r.Context(), w)
}

func HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
  return render(r, w, auth.Signup())
}

func HandleSignup(w http.ResponseWriter, r *http.Request) error {
  params := auth.SignupParams{
    Email: r.FormValue("email"),
    Password: r.FormValue("password"),
    ConfirmPassword: r.FormValue("confirmPassword"),
  }


  errors := auth.SignupErrors{}
  if ok := validate.New(&params, validate.Fields{
    "Email": validate.Rules(validate.Email),
    "Password": validate.Rules(validate.Password),
    "ConfirmPassword": validate.Rules(validate.ConfirmPassword(params.Password)),

  }).Validate(&errors); !ok {
    return render(r, w, auth.SignupForm(params, errors))

  }

  user, err := sb.Client.Auth.SignUp(r.Context(), supabase.UserCredentials{
    Email: params.Email,
    Password: params.Password,
  })


  if err != nil {
    return err
  }

  return render(r, w, auth.SignupSuccess(user.Email)) 
}


func HandleLogin(w http.ResponseWriter, r *http.Request) error {
  creds := supabase.UserCredentials{
    Email: r.FormValue("email"),
    Password: r.FormValue("password"),
  }

   resp, err := sb.Client.Auth.SignIn(r.Context(), creds)

   if err != nil {
     slog.Error("login error", "err", err)
     return render(r, w, auth.LoginForm(creds, auth.LoginErrors{
       InvalidCreds: "the creds are bad",
     }))
   }

   setAuthCookie(w, resp.AccessToken)
   http.Redirect(w, r, "/", http.StatusSeeOther)

  return nil
}


func HandleAuthCallback (w http.ResponseWriter, r *http.Request) error {
  at := r.URL.Query().Get("access_token")
  if len(at) <= 0 {
    return render(r, w, auth.CallbackScript())
  }

  return nil
}

func setAuthCookie(w http.ResponseWriter, accessToken string) {

  fmt.Printf("setting cookie %s\n", accessToken)

  cookie := &http.Cookie{
     Value: accessToken,
     Name: "at", // access token
     HttpOnly: true,
     Secure: true,
  // Safari will not set the cooking in local development if secure is true
  // Probably need https to get that? 
   }

   http.SetCookie(w, cookie);
}
