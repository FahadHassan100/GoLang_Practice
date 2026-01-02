package main

import "fmt"

//way 1
/* func add(a int, b int) int {
	return a + b
} */

//way 2
/* func add(a, b int) int {
	return a + b
} */

//__________________________________
//MULTIPLE VALUES RETURN

/* func getCars() (string, string, string) {
	return "Toyoto", "Nissan", "Honda"
} */

//__________________________________
//INHERITED FUNCTION

// way 1
/* func processIt(fn func(a int) int) {
	fmt.Println(fn(1))
} */

// way 2
func processIt() func(a int) int {
	return func(a int) int {
		return 3541
	}
}

func main() {

	//fmt.Println(add(5, 4))

	//__________________________________
	// Multi value return

	//fmt.Println(getCars())

	/* car1, car2, car3 := getCars()
	fmt.Println(car1, car2, car3) */

	/* car1, car2, _ := getCars()

	fmt.Println(car1, car2) */

	//__________________________________
	// inherited function

	//way 1
	//processIt(func(a int) int { return 1354 })

	//way 2
	fn := processIt()
	check := fn(0)
	fmt.Println(check)
}
