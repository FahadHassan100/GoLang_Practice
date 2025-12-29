package main

import "fmt"

func main() {

	/* var cars []int
	fmt.Println(cars == nil)
	//fmt.Println(cars) */

	/* var cars = make([]string, 0, 5)
	//fmt.Println(cap(cars))
	cars = append(cars, "Nissan")
	cars = append(cars, "Honda")

	fmt.Println(cars)
	fmt.Println(cap(cars)) */

	carsModels := []int{}
	carsModels = append(carsModels, 2015)
	carsModels = append(carsModels, 2016)

	fmt.Println(carsModels)
}
