// Declare a package main, which is a special package
// It tells the Go compiler that the package should compile as an executable program, not a shared library.
package main

// Import the "fmt" package, which contains functions for formatted I/O operations (like printing to the console).
import "fmt"

// The main function is the entry point of the Go program.
// When you run the program, this is the first function that gets called.
func main() {
	// Declare a variable 'a' of type 'int' (integer) and initialize it with the value 10.
	// 'int' is a data type that represents an integer.
	var a int = 10

	// Similarly, declare a variable 'b' and initialize it with the value 20.
	var b int = 20

	// Declare another variable 'sum' to store the result of the addition of 'a' and 'b'.
	var sum int = a + b

	// Use fmt.Println to print the sum to the console.
	// "Sum:" is a string literal, and 'sum' is the variable that holds the calculated sum.
	// The Println function automatically adds a new line after printing.
	fmt.Println("Sum:", sum)
}
