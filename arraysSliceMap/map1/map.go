package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234567888] = "Maria"
	aprovados[1238273282] = "Ana"
	aprovados[1287319823] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1234567888])
	delete(aprovados, 1234567888)
	fmt.Printf("%q", aprovados[1234567888])
}
