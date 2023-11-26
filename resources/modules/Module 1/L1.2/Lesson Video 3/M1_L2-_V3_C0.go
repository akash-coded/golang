// Package Declaration:
// Declares the package name 'main'. In Go, the 'main' package is special; it defines an executable program.
package main

// Import Statement:
// Imports the 'fmt' package, which contains functions for formatted I/O operations, like printing to the console.
import "fmt"

// Main Function:
// The 'main' function is the entry point of a Go program. Execution starts here when the program is run.
func main() {
	// Numeric Types Example:
	// Declares and initializes an integer variable 'age' with a value of 28.
	// 'int' is a data type for integers (whole numbers).
	var age int = 28

	// Declares and initializes an integer variable 'count' with a larger value.
	// 'int64' is a data type for larger integers (64-bit integer).
	var count int64 = 1000000000

	// Print numeric variables using fmt.Println.
	fmt.Println("Age:", age)
	fmt.Println("Count:", count)

	// String Example:
	// Declares and initializes a string variable 'str'.
	// Strings in Go are sequences of characters.
	str := "Welcome to Go programming"

	// Using fmt.Printf for formatted printing.
	// %d is the format specifier for integers, used here to print the length of 'str'.
	// %s is the format specifier for strings.
	// %T is the format specifier to print the type of a variable.
	fmt.Printf("Length is: %d", len(str))
	fmt.Printf("\nString is: %s", str)
	fmt.Printf("\nType of str is: %T", str)

	// Boolean Example:
	// Declares a boolean variable 'boolValue'.
	// Booleans in Go represent true or false values.
	var boolValue bool = false

	// Print the boolean variable using fmt.Println.
	fmt.Println("\nBoolean value:", boolValue)
}
