package service

import (
	"context"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, id string) (*models.Car, error)
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, car *models.CarRequest) (*models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (*models.Car, error)
	DeleteCar(ctx context.Context, id string) (*models.Car, error)
}

type EngineServiceInterface interface {
	GetEngById(ctx context.Context, id string) (*models.Engine, error)
	CreateEngine(ctx context.Context, engReq *models.EngineReq) (*models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engReq *models.EngineReq) (*models.Engine, error)
	DeleteEngine(ctx context.Context, id string) (*models.Engine, error)
}
