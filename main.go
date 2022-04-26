package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the temperature unit (c/f/k): ")
	var flag string
	fmt.Scanf("%s", &flag)
	fmt.Println("Enter the temperature: ")
	var temp float64
	fmt.Scanf("%s", &temp)
	if flag == "" {
		fmt.Println("Please provide a temperature unit (c/f/k).")
		return
	} else if flag == "c" {
		var tempf float64 = temp*1.8 + 32
		var tempk float64 = temp + 273.15
		fmt.Println("The temperature in Fahrenheit is: ", tempf)
		fmt.Println("The temperature in Kelvin is: ", tempk)
	} else if flag == "f" {
		var tempc float64 = (temp - 32) * 5 / 9
		var tempk float64 = (temp + 459.67) * 5 / 9
		fmt.Println("The temperature in Celsius is: ", tempc)
		fmt.Println("The temperature in Kelvin is: ", tempk)
	} else if flag == "k" {
		var tempc float64 = temp - 273.15
		var tempf float64 = 1.8*(temp-273.15) + 32
		fmt.Println("The temperature in Celsius is: ", tempc)
		fmt.Println("The temperature in Fahrenheit is: ", tempf)
	} else {
		fmt.Println("Please provide a valid temperature unit (c/f/k).")
	}
}
