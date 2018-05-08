package abstractfactory

// FamilyCar TODO
type FamilyCar struct{}

// NumDoors TODO
func (f *FamilyCar) NumDoors() int {
	return 5
}

// NumWheels TODO
func (f *FamilyCar) NumWheels() int {
	return 4
}

// NumSeats TODO
func (f *FamilyCar) NumSeats() int {
	return 4
}
