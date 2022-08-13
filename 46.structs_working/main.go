package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	// page := lastBook.pages // compile time error, pages are undefind

	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook) // +v print out both fields and their values

	aBook := book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}

	fmt.Println(aBook == lastBook)

	// create copy of struct
	myBook := aBook // myBook is a copy of book (brand new reference in memory)
	myBook.year = 2020
	fmt.Println(myBook, aBook)

	// anonymous struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's age: %d\n", diana.age)

	type Book struct {
		string // field "string" of type string is created automatically
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
}

// structs are fixed at compile time, fields cannot be added nor removed at runtime
// struct values are comparable to each other with == operator
