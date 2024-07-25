package models

const UserContextKey = "user"

type AuthenticatedUser struct {
  Email       string
  IsLoggedIn  bool
}
