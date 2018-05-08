package abstractfactory

import (
	"fmt"
	// "errors"
)
// VehicleFactory returns other Factoy Like Bike Factory CarFactory
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

// Type of Factory supported by VehicleFactory
const (
	CarFactoryType = 1
	MoterbikeFactoryType = 2
)

// BuildFactory Retuens Factory to create Car Or Moterbike
func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MoterbikeFactoryType:
		return new(MoterBikeFactory), nil

	default:
		return nil, fmt.Errorf("Factory with %d does not reconized", f)
	}
}

// CarFactory Supports cars types
const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

// CarFactory used to create LuxryCar, FamilyCar....
type CarFactory struct {}

// NewVehicle accept cartype i.e LuxuryCarType, FamilyCarType and returns LuxuryCar, FamilyCar....
func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d is not reconized", v)	
	}
}

// Sported Bike Types
const (
	SportMoterbikeType = 1
	CruiseMoterbikeType = 2
)

// MoterBikeFactory to create SportsMoterbike, CruiseMoterBike etc...
type MoterBikeFactory struct{}

// NewVehicle Return new Type MoterBike i.e SportsMoterbike, CruiseMoterBike...
func (m *MoterBikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMoterbikeType:
		return new(SportMoterbike), nil
	case CruiseMoterbikeType:
		return new(CruiseMoterbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d is not reconized", v)
	}
}