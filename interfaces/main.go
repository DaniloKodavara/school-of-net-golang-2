package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

type motorcycle struct {
	name string
}

func (m motorcycle) start() string {
	return "The motorcycle " + m.name + " has been started"
}

func (c car) start() string {
	return "The car " + c.name + " has been started"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"Gol"}
	fmt.Println(startVehicle(c))

	m := motorcycle{"XPTO"}
	fmt.Println(startVehicle(m))
}
