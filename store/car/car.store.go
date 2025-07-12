package car

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func new(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	query := `SELECT c.id, c.name, c.year, brand, c. fuel type, c.engine_id, c.price, c.created_at,
	c.updated_at, e.id, e.displacement, e.no_of_cyclinders, e.car_range FROM car c LEFT JOIN
	engine e ON c.engine_id = e.id WHERE c.id=$1`

	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&car.ID,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Engine.EngID,
		&car.Price,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Engine.EngID,
		&car.Engine.Displacement,
		&car.Engine.NoOfCylinders,
		&car.Engine.CarRange,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}
		return car, err
	}
	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string
	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c. fuel_type, c.engine_id, c.price, c.created_at, c.updated_at,
		e.id, e.displacement, e.no_of_cyclinders, e.car_range FROM car c LEFT JOIN engine e ON c.engine_id=e.id WHERE c.brand = $1`
	} else {
		query = `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car
		WHERE brand = $1`
	}
	row, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var car models.Car
		if isEngine {
			var engine models.Engine
			err := row.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngID,
				&car.Price,
				&car.CreatedAt,
				&car.UpdatedAt,
				&engine.EngID,
				&engine.Displacement,
				&engine.NoOfCylinders,
				&engine.CarRange,
			)
			if err != nil {
				return nil, err
			}
			car.Engine = engine
		} else {
			err := row.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngID,
				&car.Price,
				&car.CreatedAt,
				&car.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	var createCar models.Car
	var engineId uuid.UUID

	err := s.db.QueryRowContext(ctx, "SELECT id FROM engine WHERE id=$1", carReq.Engine.EngID).Scan(&engineId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createCar, errors.New("Engine id does not be exist in table")
		}
	}
	carId := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt
	newCar := models.Car{
		ID:        carId,
		Name:      carReq.Name,
		Year:      carReq.Year,
		Brand:     carReq.Brand,
		FuelType:  carReq.FuelType,
		Engine:    carReq.Engine,
		Price:     carReq.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// Begin the transactions
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createCar, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	query := `INSERT INTO car (id,name,year,brand,fuel_type,engine_id,price,created_at,updated_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id,name,year,brand,fuel_type,engine_id,price,created_at,updated_at`

	err = tx.QueryRowContext(ctx, query,
		newCar.ID,
		newCar.Name,
		newCar.Year,
		newCar.Brand,
		newCar.FuelType,
		newCar.Engine.EngID,
		newCar.Price,
		newCar.CreatedAt,
		newCar.UpdatedAt).Scan(
		createCar.ID,
		createCar.Name,
		createCar.Year,
		createCar.Brand,
		createCar.FuelType,
		createCar.Engine.EngID,
		createCar.Price,
		createCar.CreatedAt,
		createCar.UpdatedAt,
	)
	if err != nil {
		return createCar, err
	}
	return createCar, err

}

func (s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	var updatedCar models.Car
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updatedCar, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		err = tx.Commit()
	}()

	query := `
		UPDATED CAR
		SET name=$2,year = $3, fuel_type = $5, engine_id = $6, price  = $7, updated_at = $8
		WHERE id = $1
		RETURN id, name, year, brand, fuel_type, engine_id, price, price, created_at, updated_at
	`

	err = tx.QueryRowContext(ctx, query,
		id,
		carReq.Name,
		carReq.Year,
		carReq.Brand,
		carReq.FuelType,
		carReq.Engine.EngID,
		carReq.Price,
		time.Now()).Scan(
		updatedCar.ID,
		updatedCar.Name,
		updatedCar.Year,
		updatedCar.Brand,
		updatedCar.FuelType,
		updatedCar.Engine.EngID,
		updatedCar.Price,
		updatedCar.CreatedAt,
		updatedCar.UpdatedAt,
	)
	if err != nil {
		return updatedCar, err
	}
	return updatedCar, err

}

func (s Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	var deletedCar models.Car
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return deletedCar, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		err = tx.Commit()
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, name, brand, year, fuel_type, engine_id, price, created_at, updated_at FROM car WHERE id=$1", id).Scan(
		&deletedCar.ID,
		&deletedCar.Name,
		&deletedCar.Year,
		&deletedCar.Brand,
		&deletedCar.FuelType,
		&deletedCar.Engine.EngID,
		&deletedCar.Price,
		&deletedCar.CreatedAt,
		&deletedCar.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Car{}, errors.New("car id not Present in DB")
		}
		return models.Car{}, err
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM car WHERE id=$1", id)
	if err != nil {
		return models.Car{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Car{}, err
	}
	if rowsAffected == 0 {
		return models.Car{}, errors.New("no rows were Deleted")
	}
	return deletedCar, nil
}
