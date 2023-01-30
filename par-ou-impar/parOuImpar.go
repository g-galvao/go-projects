package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, numero := range numeros {
		if numero%2 == 0 {
			fmt.Println(numero, " é par")
		} else {
			fmt.Println(numero, " é ímpar")
		}
	}
}
