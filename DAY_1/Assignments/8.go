package main

import (
    "fmt"
    "math"
)

func isArmstrong (number int)bool {
	temp,sum,digits := 0

	temp = number
	for number != 0 {
		digits++
		number /= 10
	}
	temp = number
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit),float64(digits)))
		temp /= 10
	}
	return sum == number
}
fun main() {
	var number int

	fmt.Print("Enter the number:")
	fmt.Scan(&number)
	if isArmstrong(number) {
        fmt.Printf("%d is an Armstrong number.\n", number)
    } else {
        fmt.Printf("%d is not an Armstrong number.\n", number)
    }
}