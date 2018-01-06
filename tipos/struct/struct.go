package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// função com receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto

	// passando as chaves que serao preenchidas
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.89,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	// passando apenas os valores na ordem da struct para serem preenchidos
	produto2 := produto{"Notebook", 1.22, 0.10}
	fmt.Println(produto2.precoComDesconto())
}
