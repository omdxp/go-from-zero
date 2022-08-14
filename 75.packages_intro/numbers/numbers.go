package numbers

// it is important to name functions, variables, types with uppercase letter
// uppercase letter means that it is going to be public to other packages
// everything that starts with a lowercase letter will be accessible only in that package
func Even(n int) bool {
	return n%2 == 0
}

// this belongs only to the current package
// it is a private function
func checkPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// it is idiomatic to name the file the same name as the parent folder in case of one file
// also the package name to be the same name as the parent folder
// if you define more packages in the same folder (different source files with different packages) Go will throw a compile time error
// all files in a directory MUST belongs to the same package
// it is idiomatic to keep simple name for packages like "time" "http" and avoid using underscores
