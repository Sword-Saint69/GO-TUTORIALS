// 9.Sum of Natural Numbers
// Calculate the sum of the first N natural numbers.

package main

import "fmt"

func main () {
	var nummber int64
	fmt.Print("Enter the number: ")
	fmt.Scan(&number)

    if number < 0 {
        fmt.Println("Invalid input. Please enter a non-negative number.")
        return
    }

    sum := (number * (number + 1)) / 2
	fmt.Printf("The sum of the first %d natural numbers is: %d\n", number, sum)
}