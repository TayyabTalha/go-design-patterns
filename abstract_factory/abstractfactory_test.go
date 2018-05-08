package abstractfactory

import (
	"testing"
)

func TestMoterBikeFactory(t *testing.T) {
	moterbikeF, err := BuildFactory(MoterbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	moterbikeVehicle, err := moterbikeF.NewVehicle(SportMoterbikeType)
	if err != nil {
		t.Fatal(err)
	}

	noOfWheels := moterbikeVehicle.NumWheels()
	noOfSeats := moterbikeVehicle.NumSeats()
	if noOfWheels != 2 {
		t.Errorf("No of wheels in sport bike should be 2 and got %d", noOfWheels)
	}
	t.Logf("Moterbike has %d wheels", noOfWheels)

	if noOfSeats != 1 {
		t.Errorf("No of seats in sport bike should be 1 and got %d", noOfSeats)
	}

	sportBike, ok := moterbikeVehicle.(Moterbike)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport bike has type %d", sportBike.GetMoterbikeType())

	moterbikeVehicle, err = moterbikeF.NewVehicle(CruiseMoterbikeType)
	if err != nil {
		t.Fatal(err)
	}

	noOfWheels = moterbikeVehicle.NumWheels()
	noOfSeats = moterbikeVehicle.NumSeats()
	if noOfWheels != 2 {
		t.Errorf("No of wheels in cruise bike should be 2 and got %d", noOfWheels)
	}
	t.Logf("Moterbike has %d wheels", noOfWheels)

	if noOfSeats != 2 {
		t.Errorf("No of seats in cruise bike should be 1 and got %d", noOfSeats)
	}

	cruiseMoterbike, ok := moterbikeVehicle.(Moterbike)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Cruise bike has type %d", cruiseMoterbike.GetMoterbikeType())

	_, err = BuildFactory(20)
	if err == nil {
		t.Error("Accepting error in case of invalid type")
	}

	_, err = moterbikeF.NewVehicle(20)
	if err == nil {
		t.Error("Accepting error in case of invalid type")
	}
}

func TestCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	noOfWheels := carVehicle.NumWheels()
	noOfSeats := carVehicle.NumSeats()
	if noOfWheels != 4 {
		t.Errorf("No of wheels in Luxury car type should be 4 and got %d", noOfWheels)
	}
	t.Logf("Car has %d wheels", noOfWheels)

	if noOfSeats != 5 {
		t.Errorf("No of seats in Luxury car should be 5 and got %d", noOfSeats)
	}

	luxuryCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxurry Car have doors %d", luxuryCar.NumDoors())

	carVehicle, err = carFactory.NewVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	noOfWheels = carVehicle.NumWheels()
	noOfSeats = carVehicle.NumSeats()
	if noOfWheels != 4 {
		t.Errorf("No of wheels in family should be 2 and got %d", noOfWheels)
	}
	t.Logf("Family car has %d wheels", noOfWheels)

	if noOfSeats != 4 {
		t.Errorf("No of seats in family should be 4 and got %d", noOfSeats)
	}

	familycar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Family has door %d", familycar.NumDoors())

	_, err = carFactory.NewVehicle(20)
	if err == nil {
		t.Error("Accepting error in case of invalid type")
	}
}
