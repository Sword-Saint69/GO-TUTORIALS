package main

import "fmt"


func sumOfDigits(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10 
		sum += digit      
		num /= 10         
	}
	return sum
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)


	if number < 0 {
		number = -number
	}

	fmt.Printf("The sum of digits of %d is: %d\n", number, sumOfDigits(number))
}
