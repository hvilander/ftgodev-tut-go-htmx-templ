package handler

import (
  "net/http"
  "strings"
  "context"
  "ftgodev-tut/models"
)

func WithAccess(next http.Handler) http.Handler{
  fn := func(w http.ResponseWriter, r *http.Request){
    if strings.Contains(r.URL.Path, "/public") {
      next.ServeHTTP(w,r)
      return
    }
    // add auth user to every context
    user := models.AuthenticatedUser{ 
      Email: "agg@gmail.com",
      IsLoggedIn: true,
    }
    ctx := context.WithValue(r.Context(), models.UserContextKey, user)

    next.ServeHTTP(w,r.WithContext(ctx))
  }

  return http.HandlerFunc(fn)
}



