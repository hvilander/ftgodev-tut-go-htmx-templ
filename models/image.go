package models

import (
  "time"
  "github.com/google/uuid"
)

type ImageStatus int

const (
  ImageStatusFailed ImageStatus = iota
  ImageStatusPending 
  ImageStatusCompleted 
)

type Image struct {
  ID int `bun:"id,pk,autoincrement"`
  UserID uuid.UUID
  Status ImageStatus
  ExternalID string
  Prompt string 
  deleted bool `bun:"default:'false'"`
  DeletedAt time.Time
  CreatedAt time.Time `bun:"default:'now()'"`
  Location string
}
