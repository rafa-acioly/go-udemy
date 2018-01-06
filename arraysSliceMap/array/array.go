package main

import (
	"fmt"
)

func main() {
	// homogenea (mesmo tipo) e estatica (fixo)
	// não é possivel aumentar ou diminuir o tamanho do array após sua criação
	var notas [3]float64
	fmt.Println(notas)
	fmt.Println(len(notas)) // len() pega a quantidade de itens no array

	notas[0], notas[1], notas[2] = 7, 8.8, 9.1
	fmt.Println(notas)

	var total float64
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f", media)
}
