package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateRequest(carReq CarRequest) error {
	if err := ValidateName(carReq.Name); err != nil {
		return err
	}
	if err := ValidateBrand(carReq.Brand); err != nil {
		return err
	}
	if err := ValidateEngine(carReq.Engine); err != nil {
		return err
	}
	if err := ValidateFuelType(carReq.FuelType); err != nil {
		return err
	}
	if err := ValidatePrice(carReq.Price); err != nil {
		return err
	}
	if err := ValidateYear(carReq.Year); err != nil {
		return err
	}
	return nil
}

func ValidateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func ValidateYear(year string) error {
	if year == "" {
		return errors.New("years is required")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()
	if yearInt < 1986 || yearInt > currentYear {
		return errors.New("year must be between 1986 and current year")
	}
	return nil
}

func ValidateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func ValidateFuelType(fueltype string) error {
	fueltype = strings.ToLower(fueltype) // normalize

	validfueltype := []string{"petrol", "diesel", "cng", "electric"}

	for _, validType := range validfueltype {
		if fueltype == validType {
			return nil
		}
	}
	return errors.New("fueltype must be one of: petrol, diesel, cng, electric")
}

func ValidateEngine(engine Engine) error {
	if engine.EngID == uuid.Nil {
		return errors.New("engineId is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("noOfCylinders must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}

func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
