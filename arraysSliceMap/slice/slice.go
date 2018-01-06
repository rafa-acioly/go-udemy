package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array! slice define um pedaço de um array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:3] // novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s4)

	s5 := a2[:3]
	fmt.Println(s5)
}