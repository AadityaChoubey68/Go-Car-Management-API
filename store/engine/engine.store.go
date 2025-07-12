package engine

import (
	"context"
	"database/sql"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {

}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineReq) (models.Engine, error) {

}

func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineReq) (models.Engine, error) {

}

func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {

}
