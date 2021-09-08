package main

import "fmt"

//Um programa que converte Fahrenheit em Celsius

func main() {
	fmt.Print("Digite a temperatura em F: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9
	fmt.Println("A temperatura em C é: ", output)
}
