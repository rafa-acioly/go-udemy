package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Rafael", "Pq voce nao fala comigo?", 3)
	// fale("Nadine", "Só posso falar depois de voce!", 1)

	// go fale("Rafael", "ei...", 20)
	// go fale("Nadine", "opa...", 20)

	go fale("Maria", "entendi", 10)
	fale("João", "parabéns", 5)
}
