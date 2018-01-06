package main

import (
	"fmt"
)

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() (total float64) {
	total = 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, quantidade: 2, preco: 12.10},
			item{produtoID: 2, quantidade: 1, preco: 23.22},
			item{produtoID: 11, quantidade: 200, preco: 33.99},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
