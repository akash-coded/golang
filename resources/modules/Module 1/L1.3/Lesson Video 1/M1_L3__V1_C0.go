// Package Declaration:
// This line declares the package 'main'. In Go, the 'main' package is special as it defines a standalone executable program.
package main

// Import Statement:
// Imports the 'fmt' package, which is used for formatted input/output operations, such as printing text to the console.
import "fmt"

// Main Function:
// The 'main' function is the entry point of a Go program. Execution starts here when the program is run.
func main() {
	// If-Else Statement Example:
	// Declaring and initializing an integer variable 'age'.
	age := 25

	// Using an if-else statement to perform a conditional check on the age.
	// This block checks if the age is 18 or more.
	if age >= 18 {
		// This code executes if the condition is true.
		fmt.Println("You are an adult.")
	} else {
		// This code executes if the condition is false.
		fmt.Println("You are not yet an adult.")
	}

	// Another If-Else Example:
	// Declaring a variable 'temperature' and assigning it a value.
	temperature := 28

	// Another conditional check using the if-else statement.
	// Here, it checks if the temperature is greater than 30.
	if temperature > 30 {
		fmt.Println("It's a hot day.")
	} else {
		fmt.Println("It's not a hot day.")
	}

	// Switch Statement Example:
	// Declaring a string variable 'day' and initializing it.
	day := "Wednesday"

	// Using a switch statement to select one of many code blocks to be executed.
	// This is an alternative to the if-else-if ladder.
	switch day {
	case "Tuesday":
		fmt.Println("It's still early in the week.")
	case "Wednesday":
		fmt.Println("It's the middle of the week.")
	case "Thursday":
		fmt.Println("The weekend is approaching.")
	default:
		// The default case is executed if no other case matches.
		fmt.Println("It's the weekend!")
	}

	// Additional Example: Nested If-Else
	// Demonstrating a nested if-else statement within another if-else.
	// Checking if 'age' is greater than or equal to 13.
	if age >= 13 {
		if age < 20 {
			// This block executes if 'age' is between 13 and 19, inclusive.
			fmt.Println("You are a teenager.")
		} else {
			// This block executes if 'age' is 20 or more.
			fmt.Println("You are no longer a teenager.")
		}
	} else {
		// This block executes if 'age' is less than 13.
		fmt.Println("You are a child.")
	}
}
