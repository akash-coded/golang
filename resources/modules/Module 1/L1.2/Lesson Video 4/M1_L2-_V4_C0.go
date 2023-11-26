// Package Declaration:
// This line declares the package name 'main'. In Go, the 'main' package is for standalone executable programs.
package main

// Import Statement:
// Imports the 'fmt' package, which is used for formatted I/O operations such as printing to the console.
import "fmt"

// Constants:
// Declaring several constants. In Go, constants are variables whose values cannot be changed later.
const (
	// Basic Constants:
	Pi       = 3.14159        // A constant for the value of Pi.
	MaxValue = 100            // A constant that represents a maximum value.
	Greeting = "Hello, World" // A greeting message.
	IsReady  = true           // A boolean constant.

	// Enumerated Constants:
	// These constants represent days of the week.
	// Enumerated constants are a series of constants that are related.
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6

	// Constants with Iota:
	// 'iota' is a special identifier in Go used in constant declarations to simplify definitions of incrementing numbers.
	// Here, it's used to create a series of constants representing colors.
	Red   = iota // Red is 0
	Green        // Green is 1 (automatically incremented)
	Blue         // Blue is 2
)

// Main Function:
// The 'main' function is the entry point of a Go program.
func main() {
	// Printing Basic Constants:
	fmt.Println("Pi:", Pi)
	fmt.Println("MaxValue:", MaxValue)
	fmt.Println(Greeting)
	fmt.Println("IsReady:", IsReady)

	// Printing Enumerated Constants:
	fmt.Println("\nDays of the Week Enumerated Constants:")
	fmt.Println("Sunday:", Sunday, "Monday:", Monday, "Tuesday:", Tuesday,
		"Wednesday:", Wednesday, "Thursday:", Thursday,
		"Friday:", Friday, "Saturday:", Saturday)

	// Printing Constants with Iota:
	fmt.Println("\nColor Constants with Iota:")
	fmt.Println("Red:", Red, "Green:", Green, "Blue:", Blue)

	// Operations with Constants:
	// Demonstrating basic arithmetic operations and boolean logic with constants.
	const (
		a       = 10
		b       = 3
		sum     = a + b         // Sum of 'a' and 'b'.
		product = a * b         // Product of 'a' and 'b'.
		isTrue  = true && false // Logical AND operation on boolean values.
	)

	// Printing the results of operations with constants.
	fmt.Println("\nOperations with Constants:")
	fmt.Println("Sum of", a, "and", b, ":", sum)
	fmt.Println("Product of", a, "and", b, ":", product)
	fmt.Println("Logical AND of true and false is:", isTrue)

	// Fun Fact with Constants:
	// Calculating the circumference of a circle with a radius of 'a' using Pi.
	circumference := 2 * Pi * float64(a)
	fmt.Printf("\nFun Fact: The circumference of a circle with radius %d is %.2f\n", a, circumference)
}
