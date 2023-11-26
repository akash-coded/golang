// Main package definition
package main

// Importing the fmt package for formatted I/O operations
import "fmt"

// main function - the entry point of the program
func main() {

	// Example 1: Basic For Loop
	// This loop runs a set number of times, printing a welcome message.
	fmt.Println("Basic For Loop Example:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d: Welcome to the Program\n", i+1)
	}
	fmt.Println()

	// Example 2: For Loop with Range
	// This loop iterates over a slice (array) of strings, printing each one.
	fmt.Println("For Loop with Range Example:")
	rvariable := []string{"ABC", "Hello", "Edureka"}
	for i, j := range rvariable {
		fmt.Printf("Index %d: %s\n", i, j)
	}
	fmt.Println()

	// Example 3: Infinite Loop
	// This loop runs indefinitely. We limit it here for demonstration purposes.
	fmt.Println("Infinite Loop Example (limited to 3 iterations):")
	count := 0
	for {
		fmt.Println("Welcome to the Infinite Loop")
		count++
		if count >= 3 {
			break
		} // Break out of the loop after 3 iterations
	}
	fmt.Println()

	// Example 4: For Loop as While Loop
	// In Go, 'for' is used for 'while' type loops as well. This loop runs until a condition is met.
	fmt.Println("For Loop as While Loop Example:")
	count = 0
	for count < 3 {
		fmt.Printf("While-style iteration %d: Welcome to the Program\n", count+1)
		count++
	}
}
