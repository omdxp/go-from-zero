package main

import "fmt"

func main() {
	var employees map[string]string // initializing a map with zero value "nil"
	fmt.Printf("%#v\n", employees)

	fmt.Printf("Number of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q\n", "Omar", employees["Omar"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["random"])

	// var m1 map[[]int]string // compile time error, incomparabale map key type
	var m1 map[[3]int]string // arrays are comparable so it can be used as keys for maps
	_ = m1

	// employees["Omar"] = "Programmer" // runtime error, assignment to entry in nil map
	// we have to initialize a map to add entries

	people := map[string]float64{}
	people["Omar"] = 52331.2
	people["Jaden"] = 21.43
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.13,
		"EUR": 432.13,
		// 40:    23.31, // compile time error, all keys must have the same type as well as the value
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 500.2 // updated a value
	balances["GBP"] = 100   // add another entry
	fmt.Println(balances)

	v, ok := balances["RON"]
	if ok {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON key does not exist in the map!")
	}

	balances["RON"] = 0
	v, ok = balances["RON"]
	if ok {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON key does not exist in the map!")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD") // delete "USD" entry
	fmt.Println(balances)
}

// If an element does not exist in a map, and when accessing it it will return a zero value for that type
// keys in maps must be comparable, therefore we cannot make slices as keys in a map because they're not comparable
// to distinguish about a missing entry or a zero value, the "comma ok" idiom is used
// maps in Go have designed for fast lookup times, not fast looping
// delete() builtin function is used to delete an entry from a map with a specified key
