package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {

	car1 := Car{"Corolla", 2017, "White"}
	car2 := Car{"BMW 320i", 2017, "Black"}
	sCar := SuperCar{
		Car:    car1,
		CanFly: true,
	}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(sCar)

	fmt.Println(sCar.Name)
	fmt.Println(sCar.CanFly)
	fmt.Println(sCar.info())

}
