package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   322532532523523,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "andrei@domain.com", phone: 3244223, address: "Bucharest, Romania"}
	fmt.Println(myContact)
}

// structs can be embedded to other structs
