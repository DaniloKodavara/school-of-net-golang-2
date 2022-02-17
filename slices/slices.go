package main

import (
	"fmt"
)

func main() {
	/*	slice := make([]int, 2, 3)

		sliceTest := slice
		slice = append(slice, 1, 2)
		slice[0] = 10
		slice[1] = 8

		fmt.Println(slice)
		fmt.Println(sliceTest)*/

	/*	fmt.Printf("Tamanho: %v\n", len(slice))
		fmt.Printf("Capacidade antes: %v\n", cap(slice))

		slice = append(slice, 1)

		fmt.Printf("Capacidade depois: %v\n", cap(slice))*/

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Better 2",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0])   //Hello
	fmt.Println(sliceString[:2])  //Hello World
	fmt.Println(sliceString[1:2]) //World
	fmt.Println(sliceString[2:4]) //Much Better
	fmt.Println(sliceString[2:])  //Much Better

	fmt.Println(sliceString)
}
