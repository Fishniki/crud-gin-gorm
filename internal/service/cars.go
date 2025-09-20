package service

import (
	"context"
	"crudwebsocket/domain"
	"crudwebsocket/dto"
	"crudwebsocket/model"
	"errors"

	"github.com/google/uuid"
)

type CarsService struct {
	CarsRepository domain.CarsRepository
}

// FindById implements domain.CarsService.

func NewCars(carsRepsitory domain.CarsRepository) domain.CarsService {
	return &CarsService{
		CarsRepository: carsRepsitory,
	}
}

// Create implements domain.CarsService.
func (c *CarsService) Create(ctx context.Context, req dto.CreateCarsRequest) error {

	existingCars, err := c.CarsRepository.FindByName(ctx, req.Nama)
	if existingCars.Id != uuid.Nil {
		return errors.New("Nama sudah di gunakan")
	}

	car := model.Cars{
		Id:             uuid.New(),
		Nama:           req.Nama,
		Country:        req.Country,
		Type:           req.Type,
		Image:          req.Image,
		ProductionYear: req.ProductionYear,
	}

	if save := c.CarsRepository.Save(ctx, &car); save != nil {
		return errors.New("Gagal menyimpan data: " + err.Error())
	}

	return nil

}

// Delete implements domain.CarsService.
func (c *CarsService) Delete(ctx context.Context, id string) error {
	return c.CarsRepository.Delete(ctx, id)
}

// Index implements domain.CarsService.
func (c *CarsService) Index(ctx context.Context) ([]model.Cars, error) {
	return c.CarsRepository.FindAll(ctx)
}

// Show implements domain.CarsService.
func (c *CarsService) Show(ctx context.Context, id string) (model.Cars, error) {
	return c.CarsRepository.FindById(ctx, id)
}

// Update implements domain.CarsService.
func (c *CarsService) Update(ctx context.Context, req dto.UpdateCarsRequest) error {

	car := model.Cars{
		Id:             req.Id,
		Nama:           req.Nama,
		Country:        req.Country,
		Type:           req.Type,
		Image:          req.Image,
		ProductionYear: req.ProductionYear,
	}

	return c.CarsRepository.Update(ctx, &car)
}
