package main

import (
	"fmt"
	"time"
)

/* type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond
}

// receiver type
func (os *order) changeStatus(status string) {
	os.status = status
}

//__________________________________
//WAY 2

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
} */

//__________________________________
//WAY 3 STRUCT EMBEDDING | INHERITENCE

type customer struct {
	customerName string
	phone        string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond
	customer
}

func main() {

	/* myOrder := order{
		id:     "1",
		amount: 150,
		status: "received",
	}
	myOrder.createdAt = time.Now()

	fmt.Println("My Order", myOrder)
	//fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    200,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println("Order Two", myOrder2) */

	//__________________________________
	//WAY 2

	/* myOrder := order{
		id:     "1",
		amount: 150,
		status: "received",
	}

	gv := myOrder.amount / 0.25

	myOrder.amount = gv

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)
	*/

	//__________________________________
	//WAY 2

	/* myOrder := newOrder("3", 851, "received")

	fmt.Println(myOrder) */

	//__________________________________
	//WAY 3 INLINE STRUCTS

	/* language := struct {
		name   string
		isGood bool
	}{"goLang", true}

	fmt.Println(language) */

	//__________________________________
	//WAY 4 STRUCT EMBEDDING | INHERITENCE

	/* order4 := order{
		id:       "4",
		amount:   658,
		status:   "Paid",
		customer: customer{"fahad", "123456789"},
	} */

	customerDetails := customer{
		customerName: "Fahad",
		phone:        "132456789",
	}

	order4 := order{
		id:       "4",
		amount:   658,
		status:   "Paid",
		customer: customerDetails,
	}

	fmt.Println(order4)
}
