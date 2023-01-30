package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}

// Adiciona recebe dois inteiros e retorna a soma deles
func Adiciona(x, y int) int {
	return x + y
}

func ExampleAdicion() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
