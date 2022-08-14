package numbers

// this function checks if a number is prime or not
// it is accessbile to other packages
func IsPrime(n int) bool {
	return checkPrime(n) // this function is accessed because it belongs to the same package
}

// this function check if a number is odd
func Odd(n uint) bool {
	return n%2 != 0
}
