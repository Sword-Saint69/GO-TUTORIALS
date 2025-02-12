switch variable {
case value1:
    // Code to execute if variable == value1
case value2:
    // Code to execute if variable == value2
default:
    // Code to execute if no case matches
}

// Example usage
package main

import "fmt"

func main() {
    var grade int
    fmt.Print("Enter the grade (0-100): ")
    fmt.Scan(&grade)

    switch {
    case grade >= 90 && grade <= 100:
        fmt.Println("The Grade scored is A")
    case grade >= 80 && grade < 90:
        fmt.Println("The Grade scored is B")
    case grade >= 70 && grade < 80:
        fmt.Println("The Grade scored is C")
    case grade >= 60 && grade < 70:
        fmt.Println("The Grade scored is D")
    default:
        fmt.Println("The Grade scored is F")
    }
}

// Output:
Enter the grade (0-100)=56
The Grade scored is D

Enter the grade (0-100)=95
The Grade scored is A

Enter the grade (0-100)=101
The Grade scored is F

Enter the grade (0-100)=-10
Invalid choice. Please try again.

Enter the grade (0-100)=100
The Grade scored is A
