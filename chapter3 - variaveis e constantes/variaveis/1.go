package main

import "fmt"

func main() {
	// Uma variável é uma área de armazenagem com um tipo específico e um nome associado

	var x string           // Palavra reservada + nome da variavel + tipo
	var a string = "first" // A atribuição do valor é opcional
	var b string = "second"
	c := 10 /* Declaracao de variavel mais concisa, em que nao se precisa chamar a palavra
	reservada var e nem especificar o tipo da variavel,	o proprio compilador Go infere qual é o tipo */

	//atribui a string a x e imprime x
	x = "Hello, World!"
	fmt.Println(x)

	//atribui outra string a x e imprime x
	x = "Hello"
	fmt.Println(x)

	//concatena uma string a x (x = x + "") e imprime o novo valor de x
	x += " World"
	fmt.Println(x)

	//retorna um booleano: verifica se a é igual a b
	fmt.Println(a == b)

	fmt.Println(c)
}
