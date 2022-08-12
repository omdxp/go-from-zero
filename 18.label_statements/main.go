package main

import "fmt"

func main() {
	outer := 19 // there is no conflict between the variable and the label
	_ = outer
	people := [5]string{"Omar", "Mario", "Bea", "Philip", "Yasser"}
	friends := [2]string{"Bea", "Marry"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")
}
