// Main package definition - the starting point for a Go program.
package main

// Importing necessary packages: fmt for formatted I/O, and os for file operations.
import (
	"fmt"
	"os"
)

// main function - the entry point of the program.
func main() {
	// Example 1: Defer Statement
	demonstrateDefer()

	// Example 2: Panic and Recover
	demonstratePanicAndRecover()
}

// demonstrateDefer function shows how the defer statement is used.
func demonstrateDefer() {
	fmt.Println("Defer Statement Example:")

	// Attempt to open a file.
	file, err := os.Open("example.txt")

	// Check if file opening produced an error.
	if err != nil {
		// Print the error and return from the function.
		fmt.Println("Error opening file:", err)
		return
	}

	// Schedule the file to be closed when the function exits.
	// This ensures that the file is closed properly, even if an error occurs.
	defer file.Close()

	// Additional file operations would go here.

	fmt.Println("File opened successfully")
}

// demonstratePanicAndRecover function shows how panic and recover work.
func demonstratePanicAndRecover() {
	fmt.Println("\nPanic and Recover Example:")

	// Defer a function that will execute after the panic.
	defer func() {
		// Recover from the panic and print a recovery message.
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Print a message before the panic.
	fmt.Println("Before panic")

	// Cause a panic with a custom error message.
	panic("Something went wrong!")

	// This line will not execute because the panic stops execution flow.
	fmt.Println("After panic")
}
