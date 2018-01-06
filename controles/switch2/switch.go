package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// Omitir o valor para condiçao switch fara com que o primeiro case que seja verdadeiro seja executado
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite.")
	}
}
