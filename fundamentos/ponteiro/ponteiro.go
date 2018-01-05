package main

import "fmt"

func main() {
	i := 1

	var p *int
	p = &i // pegando o endereço da variável "i"
	*p++   // desreferenciando (pegando valor)
	i++

	// GO não tem aritmética de ponteiro
	// p++

	fmt.Println(p, *p, i, &i)

}
