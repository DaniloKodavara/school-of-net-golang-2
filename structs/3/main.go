package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}

func main() {

	/*	car := Car{"Gol", 2017, "Yellow"}

		result, _ := json.Marshal(car)

		fmt.Println(string(result))*/

	var car Car

	data := []byte(`{"Name": "Gol", "Year": 2017, "Color":"Black"}`)

	json.Unmarshal(data, &car)
	fmt.Println(car)

}
