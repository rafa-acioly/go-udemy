package main

import (
	"fmt"
)

// uint serve para numeros sem sinais (negativos)
// desta maneira impedimos que numeros negativos entrem na funcao e assim nao é preciso tratar o mesmo
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	// resultado := fatorial(-1) -- erro pois o parametro da funcao fatorial é um "uint" (numero sem sinal)
}
