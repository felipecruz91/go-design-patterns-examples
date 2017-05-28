package builder

// BuildProcess is an interface that every builder must implement
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	Builder BuildProcess
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

func (director *ManufacturingDirector) SetBuilder(b BuildProcess) {
	director.Builder = b
}

func (director *ManufacturingDirector) Construct() VehicleProduct {
	return director.Builder.SetWheels().SetSeats().SetStructure().Build()
}
