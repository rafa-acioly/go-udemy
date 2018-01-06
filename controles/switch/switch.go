package main

import "fmt"

func noteParaConceiro(n float64) string {
	var nota = int(n)

	// Por padrão o GO para a execução do switch em toda verificação (case)
	// não fazendo necessidade de usar um 'break', caso deseje continuar a validação use "fallthrough"
	switch nota {
	case 10:
		fallthrough // vai para o proximo case
	case 9:
		return "A" // retorna "A" e sai do switch pois não tem a palavra "fallthrough"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Print(noteParaConceiro(9.8))
}
