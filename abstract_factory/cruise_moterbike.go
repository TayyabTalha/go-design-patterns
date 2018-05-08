package abstractfactory

// CruiseMoterbike concreat type of Moterbike
type CruiseMoterbike struct{}

// NumSeats returns number of seats
func (c *CruiseMoterbike) NumSeats() int {
	return 2
}

// NumWheels returns number of wheels
func (c *CruiseMoterbike) NumWheels() int {
	return 2
}

// GetMoterbikeType return moterbike type i.e CruiseMoterbikeType
func (c *CruiseMoterbike) GetMoterbikeType() int {
	return CruiseMoterbikeType
}