// 4.jazzkick
// Print numbers from 1 to 100 with "jazz", "Kick", and "Jazzkick" logic.

package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Jazzkick")
		} else if i%3 == 0 {
			fmt.Println("Jazz")
		} else if i%5 == 0 {
			fmt.Println("Kick")
		} else {
			fmt.Println(i)
		}
	}
}

