package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums)) // this time the capacity is 4 so it'll not create a new backing array next time to append

	nums = append(nums, 4) // this time it will not create a new backing array
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 100)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums)) // new capacity of 16

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[:1], "X", "Y")

	fmt.Println(letters, len(letters), cap(letters))
	// fmt.Println(letters[4]) // runtime error, index out of range
	fmt.Println(letters[3:6]) // it will return [D E F] because it is from the backing array
}

// when appending to slices, if the capacity is full it'll create a new backing array
// slices stores elements in the backing array
// creating a new backing array repeadtly is not efficient at all, that's why the capacity will be larger when appending a new element
// when appending a new element tha capacity will be increased by 2 -> 8 -> 16 -> ... and so on
