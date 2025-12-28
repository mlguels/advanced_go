package main

import (
	"errors"
	"fmt"
)


type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id string
	cargo int
}

type ElectricTruck struct {
	id string
	cargo int
	battery float64
}

var (
	ErrorsNotImplemented = errors.New("not Implemented")
	ErrTruckNotFound = errors.New("could not process truck")
)

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery += 1
	return nil
}
func (e *ElectricTruck) UnloadCargo() error {
		e.cargo = 0
		e.battery = -2
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck %+v \n", truck)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}
