package main

import "encoding/json"
import "fmt"

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1899.00,
		Tags: []string{
			"Promoção",
			"Eletrônicos",
		},
	}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"caneta","preco":8.90,"tags":["papelaria", "importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
