for {
    // Code to execute
}
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter 'stop' to exit: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "stop" {
            fmt.Println("Exiting...")
            break
        }

        fmt.Println("You entered:", input)
    }
}

// Output:
