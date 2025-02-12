// 5.Factorial Using Loops
// Calculate the factorial of a number using a loop.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Factorial is not defined for negative numbers or zero.")
		return
	}

	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("The factorial of %d is: %d\n", n, factorial)
}

// Output:
// Input = 10
// Output: The factorial of 10 is: 3628800

// Input = 5
// Output: The factorial of 5 is: 120
