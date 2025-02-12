package main

import "fmt"

func celsiustofarenheit(float64) float64 {
	return (celsius * 9/5) + 32
}

func farenheittocelsius(float64) float64 {
    return (fahrenheit - 32) * 5/9
}

func main() {

	var choice int
	var temperature float64

	fmt.Println("Choose conversion type:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("Enter your choice: 1 or 2")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Enter the temperature in Celsius: ")
        fmt.Scan(&temperature)
        fmt.Println("The tempreture in farenheit is:",celsiustofarenheit(temperature))
    } else if choice == 2 {
		fmt.Print("Enter the temperature in Fahrenheit: ")
		fmt.Scan(&temperature)
		fmt.Println("The temperature in Celsius is:", farenheittocelsius(temperature))
		} else {
        fmt.Println("Invalid choice. Please try again.")
    }

}
