package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600 // even though s3 is modified, s1 and s4 are modified too
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2 // here is modified because the elements are in the backing array
	fmt.Println(slice1)
	fmt.Println(slice2)

	cars := []string{"Ferrari", "Lambo", "BMW", "Henessey", "MCLaren"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...)
	// newCars and cars don't share the same backing array
	cars[0] = "Mercedes"
	fmt.Println(cars)
	fmt.Println(newCars)

	s1 = []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) // since the share the same backing array, the capacity is the length of the backing array

	newSlice = s1[2:5] //{30, 40, 50}
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p %p\n", &s1, &newSlice) // different addresses (different variables) but the backing array is the same

	newSlice[0] = 1000
	fmt.Println("s1:", s1) // s1 have been changed

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("array's size in bytes: %d\n", unsafe.Sizeof(a))
	fmt.Printf("slice's size in bytes: %d\n", unsafe.Sizeof(s))
}

// slices stores elements in a backing array when a slice is created by slice expression
// slice expression doesn't create a new backing array (the slices are still connected by the same backing array)
// to create a new slice from existing slice we don't use slice expressions but rather the append() builtin function
// len() builtin functions return the number of elements, where cap() builtin function returns the capacity that the slice can hold (not only slices)
// the slice address is the slice header address
// arrays occupies more memory than slices
