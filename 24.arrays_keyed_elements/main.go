package main

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50} // always zero elements are used for default values
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan", // mean the last index is 5
	}
	fmt.Println(names, len(names))

	cities := [...]string{
		5:        "Berlin",
		"London", // this is the last element (index 6)
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)
}

// the unkeyed element, get its index from the last element
