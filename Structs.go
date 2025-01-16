package main

import "fmt"

// Define a struct (similar to a class)
type Car struct {
	Brand string
	Model string
	Year  int
}

// Constructor function for the Car struct
func NewCar(brand string, model string, year int) *Car {
	return &Car{Brand: brand, Model: model, Year: year}
}

// Method for the Car struct (like a class method)
func (c *Car) GetDetails() string {
	return fmt.Sprintf("Car Details: %s %s (%d)", c.Brand, c.Model, c.Year)
}

// Method to "update" the year
func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear
}

func main() {
	// Create instances of Car using the constructor
	myCar := NewCar("Tesla", "Model S", 2023)
	anotherCar := NewCar("Toyota", "Corolla", 2020)

	// Call methods on the Car instances
	fmt.Println(myCar.GetDetails())
	fmt.Println(anotherCar.GetDetails())

	// Update the year of myCar
	myCar.UpdateYear(2024)
	fmt.Println("After update:", myCar.GetDetails())
}
