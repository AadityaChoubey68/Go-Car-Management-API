package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
	"github.com/google/uuid"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("RollBack Error")
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("Tx Commit Error")
			}
		}
	}()

	err = tx.QueryRowContext(ctx, `SELECT id,displacement,no_of_cylinders,car_range FROM engine WHERE id=$1`, id).Scan(
		&engine.EngID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}

	return engine, err

}

func (e EngineStore) EngineCreated(ctx context.Context, engineReq *models.EngineReq) (models.Engine, error) {
	var createdEngine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return createdEngine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("RollBack Error")
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("Tx Commit Error")
			}
		}
	}()
	engineID := uuid.New()

	tx.ExecContext(ctx, `INSERT INTO engine (id,displacement,no_of_cylinders,car_range) VALUES ($!,$2, $3, $4)`, engineID, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange)
	if err != nil {
		return models.Engine{}, err
	}
	createdEngine = models.Engine{
		EngID:         engineID,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return createdEngine, err
}

func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineReq) (models.Engine, error) {
	engineId, err := uuid.Parse(id)
	if err != nil {
		return models.Engine{}, fmt.Errorf("invalid Engine Id %w", err)
	}
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("RollBack Error")
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("Tx Commit Error")
			}
		}
	}()

	result, err := tx.ExecContext(ctx, `
	UPDATE engine SET displacement = $1,no_of_cylinders=$2,car_range=$3 WHERE id=$4
	`, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange, engineId)

	if err != nil {
		return models.Engine{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("no rows were updated")
	}
	updatedEngine := models.Engine{
		EngID:         engineId,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}
	return updatedEngine, nil
}

func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("RollBack Error")
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("Tx Commit Error")
			}
		}
	}()

	err = tx.QueryRowContext(ctx, `SELECT id,displacement,no_of_cylinders,car_range FROM engine WHERE id=$1`, id).Scan(
		&engine.EngID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}

	result, err := tx.ExecContext(ctx, `DELETE FROM engine WHERE id=$1`, id)

	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("no rows were Deleted")
	}
	return engine, nil

}
