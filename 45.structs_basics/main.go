package main

import "fmt"

func main() {
	title1, author1, year1 := "The Divine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string // field
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Divine Comedy", "Dante Aligheri", 1320} // this is a struct literal
	fmt.Println(myBook)

	otherBook := book{title: "Animal Farm", year: 1945, author: "George Orwell"}
	_ = otherBook

	aBook := book{title: "Just a random book"} // other values will be initialized with zero values
	fmt.Printf("%#v\n", aBook)
}

// struct is bultin type, it holds fields with different types
// order matters in struct when initializing without specifying the corresponding fields
