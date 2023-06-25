package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtracao(t *testing.T) {
	sub := subtracao(100, 50)

	if sub != 50 {
		t.Error("Resultado da multiplicacao é inválido")
	}
}

func TestMultiplicacao(t *testing.T) {
	mult := multiplicacao(10, 10)

	if mult != 100 {
		t.Error("Resultado da multiplicacao é inválido")
	}
}
