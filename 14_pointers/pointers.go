package main

import "fmt"

/* func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
} */

//__________________________________
//REFERENCE

func changeNum(num *int) {
	*num = 6
	fmt.Println("In changeNum", *num)
}

func main() {

	num := 1

	//changeNum(num)

	// Pointer/Referance

	//fmt.Println("Memory address", &num)

	//fmt.Println("After changeNum in main", num)

	//__________________________________
	//REFERENCE

	changeNum(&num)
	fmt.Println("After changeNum in main", num)

}
