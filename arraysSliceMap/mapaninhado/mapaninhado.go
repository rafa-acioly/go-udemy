package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12938.11,
			"Guga pereira":   837.22,
		},
		"J": {
			"José": 123.22,
		},
		"P": {
			"Pedro junior": 223.33,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		for nome := range funcs {
			fmt.Printf("Letra %s - Nome: %s\n", letra, nome)
		}
	}
}
