package main

import "fmt"

func main() {
	friends := map[string]int{"Omar": 24, "Maria": 25}
	neighbors := friends // neighbors is not a copy of friends, it shares the same reference in memory

	friends["Maria"] = 22
	fmt.Println(neighbors)

	// to copy a map
	people := make(map[string]int) // people have a new map header
	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18
	fmt.Printf("%#v %#v\n", people, friends)

	delete(friends, "Maria")
	fmt.Println(friends)

	delete(people, "No Existing Key")
}

// when declaring a map variable, Go creates a pointer to a map header in memory, so the map references this data structure "the map header"
// delete() builtin function will not throw an error if a key does not exist in a map
