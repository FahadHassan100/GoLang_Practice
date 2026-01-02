package main

import "fmt"

func main() {

	/* m := make(map[string]string)

	m["name"] = "GOLang"
	m["email"] = "fahad@live.com"

	fmt.Println(m) */

	//__________________________________

	//ANOTHER WAY TO CREATE MAP
	/* m := map[string]int{"price": 40, "phone": 3}

	fmt.Println(m) */

	//__________________________________

	//FIND THE VALUE EXIST OR NOT
	m := map[string]int{"price": 40, "phone": 3}

	value, ok := m["pricssse"]
	fmt.Println(value)

	if !ok {
		fmt.Println("Not Exist")
	} else {
		fmt.Println(value)
	}

}
