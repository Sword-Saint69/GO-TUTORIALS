// 7.Fibonacci Series
// Generate the Fibonacci series up to a given number.

package main

import "fmt"

func main (){
	var limit int
	fmt.Println("Enter the limit for the Fibonacci series: ")
	fmt.Scan(&limit)

	if limit <= 0 {
        fmt.Println("Limit must be a positive integer.")
        return
    }

	a,b:=0,1

	fmt.Print("Fibonacci series: ")

	for <=limit {
		fmt.Print(a, " ")
        a, b = b, a+b
	}
}