// Package Declaration:
// 'package main' declares that this file belongs to the 'main' package.
// In Go, the 'main' package is special. It defines a standalone executable program, not a library.
package main

// Import Statement:
// This line imports the 'fmt' package from the Go standard library.
// The 'fmt' package is used for formatted input/output operations, such as printing text to the console.
import "fmt"

// Main Function:
// 'func main()' declares the main function, which is the entry point of the program.
// When the program is run, Go automatically calls this function first.
func main() {
	// Variable Declaration and Initialization:
	// Declaring a variable 'name' of type 'string' and initializing it with the value “Aman".
	// The type 'string' represents a sequence of characters.
	var name string = "Aman"

	// Short Variable Declaration:
	// The ':=' syntax is a shorthand for declaring and initializing variables.
	// Here, 'age' is declared as an integer with a value of 28,
	// and 'score' is declared as a float64 (a double-precision floating point) with a value of 96.5.
	age := 28
	score := 96.5

	// Printing Variables:
	// Using 'fmt.Println' to print the values of variables 'name', 'age', and 'score'.
	// This function outputs the values to the console.
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Score:", score)

	// Updating a Variable:
	// Here, the value of 'age' is updated from 28 to 27.
	// This demonstrates that variables can have their values changed after initialization.
	age = 27
	fmt.Println("Updated Age:", age)
}
