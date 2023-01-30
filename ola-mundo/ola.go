package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const russo = "russo"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaRusso = "Привет, "

func Ola(nome string, idioma string) string {

	if nome == "" {
		nome = "mundo"
	}

	return prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {

	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case russo:
		prefixo = prefixoOlaRusso
	default:
		prefixo = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
