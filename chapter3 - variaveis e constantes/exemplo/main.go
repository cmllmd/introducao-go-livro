package main

import "fmt"

func main() {
	fmt.Print("Digite um numero: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
