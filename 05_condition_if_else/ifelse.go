package main

func main() {
	//age := 18

	/* if age >= 18 {
		println("Person is an adult")
	} else {
		println("Person is not an adult")
	} */

	//___________________________________________

	/* if age >= 18 {
		println("Person is an adult")
	} else if age >= 12 {
		println("Person is teenager")
	} else {
		println("Person is not an adult")
	} */

	//___________________________________________

	/* var role = "admin"
	var hasPermission = true
	if role == "admin" || hasPermission {
		println("yes")
	} */

	//___________________________________________
	//Define vairable inside condition

	if age := 15; age >= 18 {
		println("Person is adult", age)
	} else if age >= 12 {
		println("person is teenager", age)
	}

	//Go does not have ternary operator, you will have to use normal if else
}
