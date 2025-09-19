package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCarsRequest struct {
	Nama           string    `json:"name" form:"name" validate:"required"`
	Type           string    `json:"type" form:"type" validate:"required"`
	Country        string    `json:"country" form:"country" validate:"required"`
	Image          string    //`json:"image" form:"image"`
	ProductionYear time.Time `json:"productionyear" form:"productionyear" time_format:"2006-01-02" validate:"required"`
}

type UpdateCarsRequest struct {
	Id             uuid.UUID `json:"-"`
	Nama           string    `json:"name" form:"name" validate:"required"`
	Type           string    `json:"type" form:"type" validate:"required"`
	Country        string    `json:"country" form:"country" validate:"required"`
	Image          string    //`json:"image" form:"image"`
	ProductionYear time.Time `json:"productionyear" form:"productionyear" time_format:"2006-01-02" validate:"required"`
}
