package main

import "fmt"

func main() {
	s1 := "Hi there! 😃"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello!\"")

	fmt.Println(`He says: "Hello!"`)

	s2 := `I like \n Go!`
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`Price: 100
Brand: Nike`)

	fmt.Println(`C:\Users\Cds`)
	fmt.Println("C:\\Users\\Cds")

	var s3 = "I love " + "Go " + "programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0]) // 73

	// s3[5] = 'x' // compile time error, cannot assign to s3[5]

	fmt.Printf("%s\n", s3) // %s for strings
	fmt.Printf("%q\n", s3) // %q for quoted strings
}

// strings in Go are utf-8 encoded (that's why emojis are supported)
// string literals are declared with "" (double quotes)
// raw strings (special characters do not apply) are declared with `` (backticks)
// a string is a slice of bytes
// strings are immutable in Go
