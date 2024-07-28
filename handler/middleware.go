package handler

import (
  "net/http"
  "strings"
  "context"
  "ftgodev-tut/models"
  "ftgodev-tut/pkg/sb"
)

func WithAuth(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    if strings.Contains(r.URL.Path, "/public") {
      next.ServeHTTP(w, r)
      return
    }

    user := getAuthenticatedUser(r)

    if !user.IsLoggedIn {
      http.Redirect(w, r, "/login", http.StatusSeeOther)
    }

    next.ServeHTTP(w, r)
  }

  return http.HandlerFunc(fn)
}


func WithUser(next http.Handler) http.Handler{
  fn := func(w http.ResponseWriter, r *http.Request){
    if strings.Contains(r.URL.Path, "/public") {
      next.ServeHTTP(w,r)
      return
    }

    cookie, err := r.Cookie("at") 

    if err != nil {
      // trouble getting cookie out of local storage
      next.ServeHTTP(w,r)
      return
    }

    // check cookie with supa base 
    resp, err := sb.Client.Auth.User(r.Context(), cookie.Value)

    if err != nil {
      // not auth'ed maybe?
      next.ServeHTTP(w,r)
      return
    }

    user := models.AuthenticatedUser{
      Email: resp.Email,
      IsLoggedIn: true,
    }

    ctx := context.WithValue(r.Context(), models.UserContextKey, user)
    next.ServeHTTP(w, r.WithContext(ctx))
  }

  return http.HandlerFunc(fn)
}

