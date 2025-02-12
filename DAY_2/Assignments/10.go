// 10.Palindrome String
// Check if a string is a palindrome.

package main

import (
    "fmt"
    "strings"
)

func main() {
	var input string
	fmt.println("Enter the string to check for palindrome: ")
	fmt.Scanln(&input)

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	if ispalindrome(input) {
		fmt.Println(input, "is a palindrome.")
    } else {
		fmt.Println(input, "is not a palindrome.")
	}
}
func ispalindrome(input string)bool {
	n: = strings
	for i := 0; i < n/2; i++ {
        if input[i] !=input[n-i-1] {
            return false
        }
    }
	return true
}


