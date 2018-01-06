package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	// i é o indice e numero é o valor daquele indice
	for i, numero := range numeros {
		fmt.Printf("%d -> %d\n", i, numero)
	}

	// Acessa apenas os indices
	for num := range numeros {
		fmt.Println(num)
	}

	// ignora o indice e acessa apenas os valores de cada indice
	for _, numero := range numeros {
		fmt.Printf("%d", numero)
	}
}
