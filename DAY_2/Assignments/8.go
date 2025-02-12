// 8.Multiplication Table
// Print the multiplication table of a number.

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}