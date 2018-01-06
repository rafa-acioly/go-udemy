package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"José": 12398.22,
		"Ana":  12093.22,
		"Rafa": 92832.22,
	}

	funcsESalarios["Acioy"] = 2762.99
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
