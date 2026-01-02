package main

import "fmt"

func main() {

	/* var cars []int
	fmt.Println(cars == nil)
	//fmt.Println(cars) */

	//__________________________________

	/* var cars = make([]string, 0, 5)
	//fmt.Println(cap(cars))
	cars = append(cars, "Nissan")
	cars = append(cars, "Honda")

	fmt.Println(cars)
	fmt.Println(cap(cars)) */

	//__________________________________

	/* carsModels := []int{}
	carsModels = append(carsModels, 2015)
	carsModels = append(carsModels, 2016)

	fmt.Println(carsModels) */

	//__________________________________

	//SLICE OPERATOR
	/* var carsCount = []int{1, 2, 3, 4, 5}

	fmt.Println(carsCount[0:2]) */

	//__________________________________

	//COMPARE SLICE
	/* var scar1 = []int{1, 2, 3}
	var scar2 = []int{1, 5, 3}

	fmt.Println(slices.Equal(scar1, scar2)) */

	//__________________________________

	//TOW DIMENTION SLICE
	var cars = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(cars)

}
