package main

import (
	"fmt"
	"packages/car"
)

func main() {

	car := car.Car{Name: "Gol"}
	fmt.Println(car.Start())

}
