package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só depois que foi lido")
}

func main() {
	ch := make(chan int) // canal sem buffer

	go rotina(ch)

	fmt.Println(<-ch) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-ch) // deadlocks
	fmt.Println("Fim")
}
