package domain

import (
	"context"
	"crudwebsocket/dto"
	"crudwebsocket/model"
)

type CarsRepository interface {
	FindAll(ctx context.Context)([]model.Cars, error)
	FindById(ctx context.Context, id string) (model.Cars, error)
	FindByName(ctx context.Context, name string) (model.Cars, error)
	Save(ctx context.Context, b *model.Cars) error
	Update(ctx context.Context, b *model.Cars) error
	Delete(ctx context.Context, id string) error
}

type CarsService interface {
	Index(ctx context.Context) ([]model.Cars, error)
	Show(ctx context.Context, id string) (model.Cars, error)
	Create(ctx context.Context, req dto.CreateCarsRequest) error
	Update(ctx context.Context, req dto.UpdateCarsRequest) error
	Delete(ctx context.Context, id string) error
}