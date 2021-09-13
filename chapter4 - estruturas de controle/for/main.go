package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println("i:", i)
		i++
	}
	// outra maneira:
	for y := 0; y <= 10; y++ {
		fmt.Println("y:", y)
	}
}
