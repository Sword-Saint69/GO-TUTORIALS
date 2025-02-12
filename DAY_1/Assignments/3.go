package main

import "fmt"

func main() {
	var x,y int

	fmt.Print("Enter the First Number:")
	fmt.Scan(&x)
	fmt.Print("Enter the Second Number:")
	fmt.Scan(&y)

	fmt.Printf("Printing the numbers before swapping:  x= %d y= %d \n",x,y)
    x = x + y
	y = x - y
	x = x - y

	fmt.Printf("Printing the numbers after swapping: x= %d y= %d \n",x,y)
}