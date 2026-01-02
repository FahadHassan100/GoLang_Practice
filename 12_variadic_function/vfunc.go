package main

import "fmt"

/* func sum(nums ...int) int {

	total := 0

	for _, run := range nums {
		total = total + run
	}

	return total
} */

//__________________________________
//ANY TYPE PARAMETER FUNCTION
func anyType(param ...interface{}) interface{} {

	return param
}

func main() {

	//fmt.Println(1, 2, 3.8, "Hello")

	/* result := sum(6, 8, 3, 8, 1, 3, 1)
	fmt.Println(result) */

	//__________________________________
	//ANY TYPE PARAMETER FUNCTION
	check := anyType(5, 25, true, "Fahad")

	fmt.Println(check)

}
