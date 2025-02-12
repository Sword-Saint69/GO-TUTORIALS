if condition {
    // Code to execute if condition is true
} else {
    // Code to execute if condition is false
}

// Example usage
package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println("The number is positive.")
    } else if num < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }
}