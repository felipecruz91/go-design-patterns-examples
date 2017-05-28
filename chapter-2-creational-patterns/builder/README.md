# Builder creational pattern

The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic they require.

# Objectives

A Builder design pattern tries to:

- Abstract complex creations so that object creation is separated from the object user.

- Create an object step by step by filling its fields and creating the embedded objects.

- Reuse the object creation algorithm between many objects.

# Example

The Builder design pattern has been commonly described as the relationship between a director, a few builders, and the product they build.

We will create a vehicle builder. The process of creating a vehicle (the product) is more or less the same for every kind of vehicle (car, motorbike, etc.): choose vehicle type, assemble the structure, place the wheels and then the seats.
The director is represented by the `ManufacturingDirector` type in our example.

# Acceptance criteria

- I must have a manufacturing type that constructs everything that a vehicle needs.

- When using a car builder, the `VehicleProduct` with four wheels, five seats and a structure defined as `Car` must be returned.

- When using a motorbike builder, the `VehicleProduct` with two wheels, two seats and a structure defined as `Motorbike` must be returned.

- A `VehicleProduct` built by any `BuildProcess` builder must be open to modifications.