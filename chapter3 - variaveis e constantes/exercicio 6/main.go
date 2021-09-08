package main

import "fmt"

func main() {
	fmt.Print("Digite a altura em pés: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048
	fmt.Println("A medida em metros é: ", output)
}
