package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 600.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 600.4
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivalent to p.price = 300
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	// passing values
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling changeValues():", quantity, price, name, sold)

	// passing pointers
	fmt.Println("Before calling changeValuesByPointer():", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValuesByPointer():", quantity, price, name, sold)

	// passing struct as value
	gift := Product{
		price:       100,
		productName: "Watch",
	}
	changeProduct(gift)
	fmt.Println(gift)

	// passing struct as pointer
	changeProductByPointer(&gift)
	fmt.Println(gift)

	// passing slices
	prices := []int{1, 2, 3}
	changeSlice(prices) // it is changed
	fmt.Println("prices slice after calling changeSlice():", prices)

	// passing maps
	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap) // it is changed
	fmt.Println("myMap after calling changeMap():", myMap)
}

// in Go, there is a shortcut to access the value of struct pointers without using *
// even though you see that passing variables as pointers, that does not mean that this is passing by reference
// in Go, everything is passed by value, even when passing pointers, the function will create a copy for that pointer that share the same address
// there is an exception about slices and maps, even if they're passed by value, changing their values inside the function will change also their original value
// slices and maps are already pointers for their values (slices -> backing array)
// arrays behave like strings, ints, ... functions don't change their value
// it is not good practice to pass arrays as arguments to functions, pass slices instead
// passing by values is cheaper than passing by values, pointers put pressure on garbage collector
// pass pointers to function only when it is really necessary
