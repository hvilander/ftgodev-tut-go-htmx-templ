package handler

import (
  "os"
  "log/slog"
  "net/http"
  "ftgodev-tut/view/auth"
  "ftgodev-tut/pkg/sb"
  "ftgodev-tut/db"
  "ftgodev-tut/models"
  "ftgodev-tut/pkg/kit/validate"
  "github.com/nedpals/supabase-go"
  "github.com/gorilla/sessions"
)

const (
	sessionUserKey        = "user"
	sessionAccessTokenKey = "accessToken"
)

func HandleResetPassword(w http.ResponseWriter, r *http.Request) error { 
  return render(r, w, auth.ResetPassword())
}

func HandleResetPasswordUpdate (w http.ResponseWriter, r *http.Request) error {
  return hxRedirect(w, r, "/")
}

func HandleResetPasswordCreate  (w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r)
  if err := sb.Client.Auth.ResetPasswordForEmail(r.Context(), user.Email); err != nil {
    return err
  }

  return render(r, w, auth.CheckEmailForPasswordReset(user.Email))
}

func HandleAccountSetupIndex(w http.ResponseWriter, r *http.Request) error {
  return render(r, w, auth.AccountSetup())
}

func HandleAccountSetup(w http.ResponseWriter, r *http.Request) error {
  params := auth.AccountSetupParams{
    Username: r.FormValue("username"),
  }
  var errors auth.AccountSetupErrors
  ok := validate.New(&params, validate.Fields{
    "Username": validate.Rules(validate.Min(2), validate.Max(50)),
  }).Validate(&errors)

  if !ok {
    return render(r,w, auth.AccountSetupForm(params, errors))
  }

  user := getAuthenticatedUser(r)
  account := models.Account{
    UserID: user.ID,
    Username: params.Username,
  }

  if err := db.CreateAccount(&account); err != nil {
    return err
  }

  return hxRedirect(w,r, "/")
}


func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
  return auth.Login().Render(r.Context(), w)
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

   if err := setAuthSession(r, w, resp.AccessToken); err != nil {
     return err
   }

   return hxRedirect(w, r, "/")
}

func HandleLoginWithGoogle(w http.ResponseWriter, r *http.Request) error {
  resp, err := sb.Client.Auth.SignInWithProvider(supabase.ProviderSignInOptions{
    Provider: "google",
    RedirectTo: "http://localhost:3000/auth/callback",
  })

  if err != nil {
    return err
  }


  http.Redirect(w, r, resp.URL, http.StatusSeeOther)
  return nil
}

func HandleLogout(w http.ResponseWriter, r *http.Request) error {
  store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
  session, _ := store.Get(r, sessionUserKey)
  session.Values[sessionAccessTokenKey] = "" 
  
  if err := session.Save(r,w); err != nil {
    return err
  }

  return hxRedirect(w, r, "/login")
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

func HandleAuthCallback (w http.ResponseWriter, r *http.Request) error {
  at := r.URL.Query().Get("access_token")
  if len(at) <= 0 {
    return render(r, w, auth.CallbackScript())
  }

  if err := setAuthSession(r, w, at); err != nil {
    return err
  }

  http.Redirect(w, r, "/", http.StatusSeeOther)

  return nil
}

func setAuthSession(r *http.Request, w http.ResponseWriter, accessToken string) error {
  store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
  session, _ := store.Get(r, sessionUserKey)
  session.Values[sessionAccessTokenKey] = accessToken

  return session.Save(r,w)
}


