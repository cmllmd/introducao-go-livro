package main

import "fmt"

func main() {
	var number int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("O numero é par")
	}
	if number%2 != 0 {
		fmt.Println("O numero é impar")
	}
}
