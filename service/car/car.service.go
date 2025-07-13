package car

import (
	"context"
	"fmt"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
	"github.com/AadityaChoubey68/Go-Car-Management-API/store"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) GetCarById(ctx context.Context, id string) (*models.Car, error) {
	car, err := s.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	cars, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *CarService) CreateCar(ctx context.Context, car *models.CarRequest) (*models.Car, error) {
	if err := models.ValidateRequest(*car); err != nil {
		return nil, err
	}

	createdCar, err := s.store.CreateCar(ctx, car)
	if err != nil {
		return nil, err
	}
	return &createdCar, err
}

func (s *CarService) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (*models.Car, error) {
	if err := models.ValidateRequest(*carReq); err != nil {
		return nil, fmt.Errorf("error while validating update req body")
	}
	updatedCar, err := s.store.UpdateCar(ctx, id, carReq)
	if err != nil {
		return nil, fmt.Errorf("error returned from store update car func : %w", err)
	}
	return &updatedCar, err
}

func (s *CarService) DeleteCar(ctx context.Context, id string) (*models.Car, error) {
	deletedCar, err := s.store.DeleteCar(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedCar, nil
}
