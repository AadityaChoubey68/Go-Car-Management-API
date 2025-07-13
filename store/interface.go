package store

import (
	"context"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
)

type CarStoreInterface interface {
	GetCarById(ctx context.Context, id string) (models.Car, error)
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error)
	DeleteCar(ctx context.Context, id string) (models.Car, error)
}

type EngineStoreInterface interface {
	GetEngineById(ctx context.Context, id string) (models.Engine, error)
	EngineCreated(ctx context.Context, engineReq *models.EngineReq) (models.Engine, error)
	EngineUpdate(ctx context.Context, id string, engineReq *models.EngineReq) (models.Engine, error)
	EngineDelete(ctx context.Context, id string) (models.Engine, error)
}
