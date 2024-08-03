package handler


import (
  "os"
  "database/sql"
  "errors"
  "net/http"
  "strings"
  "context"
  "ftgodev-tut/models"
  "ftgodev-tut/pkg/sb"
  "ftgodev-tut/db"

  "github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func WithAuth(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    if strings.Contains(r.URL.Path, "/public") {
      next.ServeHTTP(w, r)
      return
    }

    user := getAuthenticatedUser(r)

    if !user.IsLoggedIn {
      // usability setup a a redirect after auth
      //path := r.URL.Path
      //http.Redirect(w, r, "/login?to=" + path, http.StatusSeeOther)
      // would need to set a cookie, then read that in the login handler
      http.Redirect(w, r, "/login", http.StatusSeeOther)
    }

    next.ServeHTTP(w, r)
  }

  return http.HandlerFunc(fn)
}

func WithAccountSetup(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    user := getAuthenticatedUser(r)
    account, err := db.GetAccountByUserId(user.ID)

    // the user has not setup acct yet
    // redirect to /account/setup
    if err != nil {
      if errors.Is(err, sql.ErrNoRows) {
        http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
        return
      }
      next.ServeHTTP(w, r)
      return
    }

      user.Account = account
      ctx := context.WithValue(r.Context(), models.UserContextKey, user)
      next.ServeHTTP(w, r.WithContext(ctx))

  }
  return http.HandlerFunc(fn)
}


func WithUser(next http.Handler) http.Handler{
  fn := func(w http.ResponseWriter, r *http.Request){
    if strings.Contains(r.URL.Path, "/public") {
      next.ServeHTTP(w,r)
      return
    }

		store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
		session, err := store.Get(r, sessionUserKey)

    if err != nil {
      next.ServeHTTP(w,r)
      return
    }

    accessToken := session.Values[sessionAccessTokenKey]
    if accessToken == nil {
      next.ServeHTTP(w, r)
      return
    }

		resp, err := sb.Client.Auth.User(r.Context(), accessToken.(string))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user := models.AuthenticatedUser{
			ID:          uuid.MustParse(resp.ID),
			Email:       resp.Email,
			IsLoggedIn:    true,
			//AccessToken: accessToken.(string),
		}

		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

  return http.HandlerFunc(fn)
}

