package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.5
	y := 2
	fmt.Println(x / (float64(y)))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("verdadeiro")
	}
}
