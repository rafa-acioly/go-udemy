package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

// um ponteiro é representado por um *
func inc2(n *int) {
	// * é usado para pegar o valor de um determinado ponteiro
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n) // passando o endereço de memoria de n
	fmt.Println(n)
}
