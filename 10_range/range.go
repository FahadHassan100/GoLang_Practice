package main

import "fmt"

func main() {

	//USE RANGE FOR ITERATION
	//Normal For loop
	/* cars := []int{5, 6, 7, 8}

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	} */

	//__________________________________

	// RANGE USE IN LOOP
	/* cars := []int{5, 6, 7, 8}

	//underscore is index of slice
	for _, car := range cars {
		fmt.Println(car)
	} */

	//__________________________________

	// USE MAP IN RANGE LOOP

	makeMap := map[string]string{"firstname": "Talha", "lastname": "Anjum"}

	for key, value := range makeMap {
		fmt.Println(key, value)
	}

}
