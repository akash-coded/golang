// Main package definition. Every Go program starts with a package declaration.
package main

// Importing the fmt package, which contains functions for formatted I/O.
import "fmt"

// main function - the entry point of our program.
func main() {

	// Example 1: Break in Loop
	// Demonstrates the use of 'break' to exit a loop prematurely.
	fmt.Println("Break Example:")
	for a := 1; a <= 6; a++ {
		// If 'a' equals 4, exit the loop immediately.
		if a == 4 {
			fmt.Println("Breaking out of the loop when a equals", a)
			break
		}
		// Print the value of 'a' in each iteration.
		fmt.Println("Value of a:", a)
	}
	fmt.Println()

	// Example 2: Continue in Loop
	// Demonstrates the use of 'continue' to skip the current loop iteration.
	fmt.Println("Continue Example:")
	for a := 1; a <= 6; a++ {
		// If 'a' equals 4, skip this iteration and continue with the next one.
		if a == 4 {
			fmt.Println("Skipping the value", a)
			continue
		}
		// Print the value of 'a' except when it is 4.
		fmt.Println("Value of a:", a)
	}
}
