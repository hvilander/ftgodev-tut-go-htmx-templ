package handler

import (
  "log/slog"
  "net/http"
  "ftgodev-tut/models"
  "github.com/a-h/templ"
)

func getAuthenticatedUser(r *http.Request) models.AuthenticatedUser {
  user, ok := r.Context().Value(models.UserContextKey).(models.AuthenticatedUser)
  if !ok {
    return models.AuthenticatedUser{}
  }
  return user

}

func MakeHandler(
  h func( http.ResponseWriter, *http.Request) error) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if err := h(w, r); err != nil {
      slog.Error("internal server error", "err", err, "path", r.URL.Path)
    }
  }
}

func render(r *http.Request, w http.ResponseWriter, component templ.Component) error {
  return component.Render(r.Context(),w)
}

func hxRedirect(w http.ResponseWriter, r *http.Request, to string) error {
  if len(r.Header.Get("HX-Request")) > 0 {
    w.Header().Set("HX-Redirect", to)
    w.WriteHeader(http.StatusSeeOther)
    return nil
  }
  
  http.Redirect(w, r, to, http.StatusSeeOther)
  return nil
}
