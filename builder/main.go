package builder

type BuilderProcess interface {
	SetWheels() BuilderProcess
	SetSeats() BuilderProcess
	SetStructure() BuilderProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuilderProcess
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder whold set Builder
func (m *ManufacturingDirector) SetBuilder(b BuilderProcess) {
	m.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuilderProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuilderProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuilderProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetSeats() BuilderProcess {
	b.v.Seats = 1
	return b
}

func (b *BikeBuilder) SetWheels() BuilderProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuilderProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}
