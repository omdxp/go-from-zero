package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c car) changeCar1(newBrand string, newPrice int) {
	// in fact, the compiler warns about the ineffective assignment to fields
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

type distance *int

// func (d *distance) m1() { // compile time error, invalid receiver (pointer or interface type)
// 	fmt.Println("just a message")
// }

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "BMW", 50000) // no change here (the function worked for a copy)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21000) // no change here (same pass by value mechanism)
	fmt.Println(myCar)

	(&myCar).changeCar2("Mercedes", 50000) // there is change (because of the pointer recever)
	fmt.Println(myCar)

	myCar.changeCar2("Seat", 25000)
	fmt.Println(myCar)

	var yourCar *car = &myCar
	fmt.Println(*yourCar)

	// valid ways to call methods
	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar) // this will also change myCar variable (pointers are here)

	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar)

	fmt.Println(myCar)
}

// Go helps with the shortcut if the receiver is a pointer, the address of operator can be omitted
// the compile will add & for us
// if a method takes a pointer receiver it's good to convert all the methods to take a pointer receiver as well
// pointer receivers are used when the receiver needs to be modified or avoid copying large data that is automatically copied when passing values
// if not, values receivers can be used
// methods are created only for value types, not pointer nor interface types
