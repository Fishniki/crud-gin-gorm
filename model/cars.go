package model


import (
	"time"

	"github.com/google/uuid"
)

type Cars struct {
	Id      uuid.UUID `gorm:"primaryKey"`
	Nama    string `gorm:"not null"`
	Type    string `gorm:"not null"`
	Country string `gorm:"not null"`
	Image string 
	ProductionYear time.Time `gorm:"type:date;not null"`
}