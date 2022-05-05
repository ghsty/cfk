package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the temperature unit (c/f/k): ")
	var flag string
	fmt.Scan(&flag)

	if flag == "" {
		fmt.Println("Please provide a temperature unit (c/f/k).")
		return

	} else if flag == "c" {
		fmt.Println("Enter the temperature: ")
		var temp float64
		fmt.Scan(&temp)
		var tempf float64 = temp*1.8 + 32
		var tempk float64 = temp + 273.15
		fmt.Println("The temperature you entered is: ", temp, "C")
		fmt.Println("The temperature in Fahrenheit is: ", tempf, "F")
		fmt.Println("The temperature in Kelvin is: ", tempk, "K")

	} else if flag == "f" {
		fmt.Println("Enter the temperature: ")
		var temp float64
		fmt.Scan(&temp)
		var tempc float64 = (temp - 32) * 0.5556
		var tempk float64 = (temp + 459.67) / 1.8
		fmt.Println("The temperature you entered is: ", temp, "F")
		fmt.Println("The temperature in Celsius is: ", tempc, "C")
		fmt.Println("The temperature in Kelvin is: ", tempk, "K")

	} else if flag == "k" {
		fmt.Println("Enter the temperature: ")
		var temp float64
		fmt.Scan(&temp)
		var tempc float64 = temp - 273.15
		var tempf float64 = 1.8*(temp-273.15) + 32
		fmt.Println("The temperature you entered is: ", temp, "K")
		fmt.Println("The temperature in Celsius is: ", tempc, "C")
		fmt.Println("The temperature in Fahrenheit is: ", tempf, "F")

	} else {
		fmt.Println("Please provide a valid temperature unit (c/f/k).")
	}
}
