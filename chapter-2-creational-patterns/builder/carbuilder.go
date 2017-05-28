package builder

type CarBuilder struct {
	vehicleProduct VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicleProduct.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicleProduct.Seats = 5
	return c

}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.vehicleProduct.Structure = "Car"
	return c

}

func (c *CarBuilder) Build() VehicleProduct {
	return c.vehicleProduct
}
