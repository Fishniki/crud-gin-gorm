package repository

import (
	"context"
	"crudwebsocket/domain"
	"crudwebsocket/model"

	"gorm.io/gorm"
)

type carsRepository struct {
	db *gorm.DB
}

// FindById implements domain.CarsRepository.
func (r *carsRepository) FindById(ctx context.Context, id string) (car model.Cars, err error) {
	if err := r.db.WithContext(ctx).First(&car, "id = ?", id).Error; err != nil {
		return car, err
	}

	return car, nil
}

// FindByName implements domain.CarsRepository.
func (r *carsRepository) FindByName(ctx context.Context, name string) (car model.Cars, err error) {
	
	if err := r.db.WithContext(ctx).Where("nama = ?", name).First(&car).Error; err != nil{
		return car, err
	}

	return car, nil

}

func NewCars(con *gorm.DB) domain.CarsRepository {
	return &carsRepository{
		db: con,
	}
}

// Delete implements domain.CarsRepository.
func (r *carsRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Cars{}, "id = ?", id).Error
}

// FindAll implements domain.CarsRepository.
func (r *carsRepository) FindAll(ctx context.Context) (cars []model.Cars, err error) {

	if err := r.db.WithContext(ctx).Find(&cars).Error; err != nil {
		return []model.Cars{}, err
	}

	return cars, nil

}

// FindById implements domain.CarsRepository.
// func (r *carsRepository) FindById(ctx context.Context, id uint) (car model.Cars, err error) {

// 	if err := r.db.WithContext(ctx).First(&car, "id = ?", id).Error; err != nil {
// 		return car, err
// 	}

// 	return car, nil

// }

// Save implements domain.CarsRepository.
func (r *carsRepository) Save(ctx context.Context, b *model.Cars) error {
	return r.db.WithContext(ctx).Save(b).Error
}

// Update implements domain.CarsRepository.
func (r *carsRepository) Update(ctx context.Context, b *model.Cars) error {

	return r.db.WithContext(ctx).Save(b).Error

}
