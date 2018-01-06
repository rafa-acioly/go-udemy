package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10) // cria um slice com 10 posições
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // cria um slice com 10 posições que aponta para um array que possui 20 posições
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // preenche completamente o slice
	fmt.Println(s, len(s), cap(s))

	// quando um slice ultrapassa a quantidade de posições do array que ele referencia, Go automaticamente cria um novo array com o mesmo
	// valor inicial definido (20) e inclui no mesmo slice ficando assim com 40 posições no array com um slice apontando para 21 posições.
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
