package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (somente positivo)... unit8 unit16 unit32 unit64
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	//com sinal... unit8 unit16 unit32 unit64
	i1 := math.MaxInt64
	fmt.Println("o valor máximo do int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("o rune é", reflect.TypeOf(i2))

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("o tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "hello world"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `olá
world,
my
name
is
gabriel`
	fmt.Println("o tamanho da string é", len(s2))
	//char?
	//var x char 'b'
	char := 'a'
	fmt.Println("o tipo de char é", reflect.TypeOf(char))
	fmt.Println("a representa", char)

}
