package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1 // envia dados para o canal (escrita)
	<-ch    // lê o dado do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
