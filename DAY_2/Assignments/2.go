// Grade Calculator
// Assign grades (A, B, C, D, F) based on a score.

package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter the Marks Scored:")
	fmt.Scan(&score)

	switch {
	case score >= 90 && score <= 100:
		fmt.Println("The Grade scored is A+")
	case score >= 80 && score < 90:
		fmt.Println("The Grade scored is A")
	case score >= 70 && score < 80:
		fmt.Println("The Grade scored is B")
	case score >= 60 && score < 70:
		fmt.Println("The Grade scored is C")
	case score >= 50 && score < 60:
		fmt.Println("The Grade scored is D")
	default:
		fmt.Println("The Grade scored is F")
	}
}
