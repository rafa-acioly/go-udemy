package main

import (
	"fmt"
	"time"
)

// Channel (Canal) - é a forma de comunicação entre goroutines
// channel é um tipo como qualquer outro (float, boolean, string e etc)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)
	go doisTresQuatroVezes(2, ch)
	fmt.Println("A")

	a, b := <-ch, <-ch // recebe dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-ch)
}
