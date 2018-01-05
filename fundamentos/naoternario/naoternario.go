package main

import (
	"fmt"
)

// Não existe operador ternario em GO

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
