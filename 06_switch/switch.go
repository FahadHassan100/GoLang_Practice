package main

func main() {

	/* i := 2

	switch i {
	case 1:
		println("This is case one")

	case 2:
		println("This is case Two")

	case 3:
		println("This is case three")
	} */

	//_______________________________________
	//MULTIPLE CONDITION SWITCH

	/* switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("it's weekend")
	default:
		println("it's working day")
	} */

	//_______________________________________
	//TYPE SWITCH

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("its an integer")
		case string:
			println("its a string")
		case bool:
			println("its bool")
		default:
			println("other")
		}
	}

	whoAmI("This is Fahad")

}
