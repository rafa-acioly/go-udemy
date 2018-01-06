package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 10, 20) // cria um slice que aponta para 10 posições de um array com 20 posições
	s2 := append(s1, 1, 2, 3) // adiciona mais 3 posições ao slice
	fmt.Println(s1, s2)

	s1[0] = 7           // altera a posição 0 do array
	fmt.Println(s1, s2) // o slice s1 e s2 mudam automaticamente o valor alterado no array
}
