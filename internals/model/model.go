package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Weight struct {
	gorm.Model           //Adds some metadata to fields to the table
	ID         uuid.UUID `gorm:"type:uuid"`
	Weight     float32
}
