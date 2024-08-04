package db

import (
  "github.com/google/uuid"
  "context"
  "ftgodev-tut/models"
)

func GetAccountByUserId (id uuid.UUID) (models.Account, error) {
  var account models.Account
  err := Bun.NewSelect().
    Model(&account).
    Where("user_id = ?", id ).
    Scan(context.Background())

  return account, err
}


func CreateAccount(account *models.Account) error {
  _, err := Bun.NewInsert().
    Model(account).
    Exec(context.Background()) // contriversal maybe don't do this?
  return err
}

func UpdateAccount(account *models.Account) error {
  _, err := Bun.NewUpdate().
    Model(account).
    WherePK().
    Exec(context.Background())

  return err
}
  
