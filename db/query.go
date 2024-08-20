package db

import (
  "github.com/google/uuid"
  "context"
  "ftgodev-tut/models"
)

func CreateImage(image *models.Image) error {
  _, err := Bun.NewInsert().
    Model(image).
    Exec(context.Background())
  return err
}

func GetImageByID(id int) (models.Image, error) {
  var image models.Image
  err := Bun.NewSelect().
    Model(&image).
    Where("id= ? ", id).
    Scan(context.Background())
  return image, err
}

func GetImagesByBatchID(batchID uuid.UUID) ([]models.Image, error) {
  var images []models.Image
  err := Bun.NewSelect().
    Model(&images).
    Where("batch_id = ? ", batchID).
    Order("created_at desc").
    Scan(context.Background())
  return images, err
}

func GetImagesByUserID(userID uuid.UUID) ([]models.Image, error) {
  var images []models.Image
  err := Bun.NewSelect().
    Model(&images).
    Where("deleted = ? ", false).
    Where("user_id = ? ", userID).
    Order("created_at desc").
    Scan(context.Background())
  return images, err
}
func UpdateImage(image *models.Image) error {
  _, err := Bun.NewUpdate().
    Model(image).
    WherePK().
    Exec(context.Background())

  return err
}

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
  
