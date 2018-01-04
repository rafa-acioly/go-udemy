package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float64
	y := 2   // int

	// fmt.Println(x / y) -- erro pois as variaveis não são do mesmo tipo
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97)) // a função string converte o numero para o valor Unicode, neste exemplo é a letra "a"

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// Qualquer numero passado para o ParseBool que não seja 1 será falso
	// b, _ := strconv.ParseBool("true") -- verdadeiro
	// b, _ := strconv.ParseBool("1") -- verdadeiro
	// b, _ := strconv.ParseBool("20") -- falso
	// b, _ := strconv.ParseBool("-12938") -- falso
	// b, _ := strconv.ParseBool("asiudhsa") -- falso
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	}
}
