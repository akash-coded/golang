// Package main - the starting point for every Go program.
package main

// Importing the fmt package, which is used for formatted I/O operations.
import "fmt"

// main function - the entry point of the program.
func main() {
	// Initializing the variable 'i' to 0.
	i := 0

	// Label 'start' - used as a target for the goto statement.
start:

	// Check if 'i' is less than 5. If so, execute the block.
	if i < 5 {
		// Print the current value of 'i'.
		fmt.Println(i)

		// Increment 'i' by 1.
		i++

		// Use the goto statement to jump to the 'start' label.
		// This creates a loop-like behavior.
		goto start
	}

	// After 'i' reaches 5, the program exits the if block and continues here.

	// Printing a completion message.
	fmt.Println("Loop completed using 'goto'.")
}
