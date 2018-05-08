package abstractfactory

// SportMoterbike Concreate Implemenation of SportBike
type SportMoterbike struct {}

// NumWheels returns total number of wheels of SportMoterbike
func (m *SportMoterbike) NumWheels() int {
	return 2
}

// NumSeats returns total num of seats in SportMoterbike
func (m *SportMoterbike) NumSeats() int {
	return 1
}

// GetMoterbikeType Returns MoterbikeType i.e SportMoterbikeType
func (m *SportMoterbike) GetMoterbikeType() int {
	return SportMoterbikeType
} 