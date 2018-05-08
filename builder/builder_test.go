package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Seats != 5 {
		t.Errorf("Seats must be 4 and they were %d\n", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure must be 'Car' and that is %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats must be 5 and they where %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.Build()

	if bike.Wheels != 2 {
		t.Errorf("Wheel must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure Must be 'Bike' and that is %s\n", bike.Structure)
	}

	if bike.Seats != 1 {
		t.Errorf("Seats must be 1 and that are %d\n", bike.Seats)
	}

}
