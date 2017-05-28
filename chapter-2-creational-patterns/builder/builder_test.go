package builder

import "testing"

func TestBuilderPattern_Car(t *testing.T) {
	// Arrange
	director := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	director.SetBuilder(carBuilder)
	// Act
	car := director.Construct()
	// Assert
	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}
}

func TestBuilderPattern_Bike(t *testing.T) {
	//Arrange
	director := ManufacturingDirector{}
	bikeBuilder := &BikeBuilder{}
	director.SetBuilder(bikeBuilder)
	//Act
	bike := director.Construct()
	//Assert
	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("Structure on a bike must be 'Motorbike' and was %s\n", bike.Structure)
	}

	if bike.Seats != 1 {
		t.Errorf("Seats on a bike must be 1 and they were %d\n", bike.Seats)
	}
}
