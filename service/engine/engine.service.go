package engine

import (
	"context"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
	"github.com/AadityaChoubey68/Go-Car-Management-API/store"
	"go.opentelemetry.io/otel"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (e *EngineService) GetEngById(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "GetEngById-Service")
	defer span.End()
	engine, err := e.store.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (e *EngineService) CreateEngine(ctx context.Context, engReq *models.EngineReq) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "CreateEngine-Service")
	defer span.End()
	if err := models.ValidateEngineReq(*engReq); err != nil {
		return nil, err
	}

	createdEng, err := e.store.EngineCreated(ctx, engReq)
	if err != nil {
		return nil, err
	}
	return &createdEng, nil
}

func (e *EngineService) UpdateEngine(ctx context.Context, id string, engReq *models.EngineReq) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Service")
	defer span.End()
	if err := models.ValidateEngineReq(*engReq); err != nil {
		return nil, err
	}

	updatedEng, err := e.store.EngineUpdate(ctx, id, engReq)
	if err != nil {
		return nil, err
	}
	return &updatedEng, nil
}

func (e *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Service")
	defer span.End()
	deletedEng, err := e.store.EngineDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedEng, nil
}
