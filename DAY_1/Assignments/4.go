package main

import "fmt"

func main() {
	const pi float64 = 3.14
	var r float32

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&r)

	area := pi * r * r
	fmt.Printf("The area of the circle is: %f \n", area)
}
