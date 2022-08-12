package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a[1:3] // from index 1 to index 2 (3 is excluded)
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:]) // same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // same as s1[0:3]
	fmt.Println(s1[:])  // same as s1[0:len(s1)]
	// fmt.Println(s1[:100]) // runtime error, slice bounds out of range

	s1 = append(s1[:4], 100) // it will add 100 after the element of index 3
	fmt.Println(s1)

	s1 = append(s1[:4], 200) // it will overwrite the last element
	fmt.Println(s1)
}

// slicing a slice returns a slicea (a[start:stop])
// a missing start index defaults to 0, and a missing stop index defaults to the last one
