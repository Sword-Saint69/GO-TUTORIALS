package main

import "fmt"


func reverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10  
		reversed = reversed*10 + digit 
		num /= 10  
	}
	return reversed
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	isNegative := number < 0
	if isNegative {
		number = -number 
	}

	reversed := reverseNumber(number)

	if isNegative {
		reversed = -reversed
	}

	fmt.Printf("The reversed number is: %d\n", reversed)
}
