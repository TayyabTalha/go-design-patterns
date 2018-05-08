package abstractfactory

// LuxuryCar concreate type of car
type LuxuryCar struct{}

// NumDoors returns num of door
func (l *LuxuryCar) NumDoors() int {
	return 4
}

// NumSeats TODO
func (l *LuxuryCar) NumSeats() int {
	return 5
}

// NumWheels TODO
func (l *LuxuryCar) NumWheels() int {
	return 4
}
