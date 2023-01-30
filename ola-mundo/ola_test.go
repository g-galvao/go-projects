package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Gabs", "")
		esperado := "Olá, Gabs"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Amelie", "francês")
		esperado := "Bonjour, Amelie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em russo", func(t *testing.T) {
		resultado := Ola("Alexander", "russo")
		esperado := "Привет, Alexander"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}
