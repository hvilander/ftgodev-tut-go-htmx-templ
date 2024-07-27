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

    user := models.AuthenticatedUser{}
    ctx := context.WithValue(r.Context(), models.UserContextKey, user)

    next.ServeHTTP(w,r.WithContext(ctx))
  }

  return http.HandlerFunc(fn)
}



