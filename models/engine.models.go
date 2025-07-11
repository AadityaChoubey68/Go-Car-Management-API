package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngID         uuid.UUID `json:"eng_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_range"`
}
type EngineReq struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"no_of_cylinders"`
	CarRange      int64 `json:"car_range"`
}

func ValidateEngineReq(engReq EngineReq) error {
	if err := ValidateDisplacement(engReq.Displacement); err != nil {
		return err
	}
	if err := ValidateNoOfCylinders(engReq.NoOfCylinders); err != nil {
		return err
	}
	if err := ValidateCarRange(engReq.CarRange); err != nil {
		return err
	}
	return nil
}

func ValidateDisplacement(disp int64) error {
	if disp <= 0 {
		return errors.New("disp must be greater than zero")
	}
	return nil
}

func ValidateNoOfCylinders(noOfCylinder int64) error {
	if noOfCylinder <= 0 {
		return errors.New("noOfCylinder must be greater than zero")
	}
	return nil
}

func ValidateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}
