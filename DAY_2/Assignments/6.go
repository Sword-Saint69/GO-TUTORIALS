// 6.Prime Number Checker
// Check if a number is prime.

package main

import "fmt"

func main (){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num <= 1 {
		fmt.Println(num, "Enter a non negative number")
	} else {
		isPrime := true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
                isPrime = false
                break
            }
		}
	}
	if isPrime {
        fmt.Println(num, "is a prime number.")
    } else {
        fmt.Println(num, "is not a prime number.")
    }
}

// Output:
// The number is 13
// 13 is a prime number.

// The number is 12
// 12 is not a prime number.

// The number is 1
// 1 is not a prime number.

// The number is -1
// -1 is not a prime number.