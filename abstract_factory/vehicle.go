package abstractfactory

// Vehicle All vehicles have provide these methods
type Vehicle interface {
	NumWheels() int
	NumSeats() int
}