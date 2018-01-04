package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516
	// fmt.Println("O valor de x é:" + x) -- erro pois "x" é float64 e não pode ser exibido como string
	xs := fmt.Sprint(x) // converte o valor para string
	fmt.Println("O valor de x é:" + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é: %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"

	// %d - inteiro
	// %f - float
	// %.2f - float com 2 casas decimais
	// %t - boolean
	// %s - string
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// %v - genérico (imprimi todo os tipos)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
