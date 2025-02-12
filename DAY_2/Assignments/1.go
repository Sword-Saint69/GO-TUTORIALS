//Even or Odd
//Check if a number is even or odd.

package main

import "fmt"

func main(){
	var num int

	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	if num % 2 == 0 {
		fmt.Printf("The number must be even: %v\n", num)
	}else{
		fmt.Println("The number must be odd: %v\n", num)
	}
}