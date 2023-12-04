package main

import "testing"


func TestSoma(t *testing.T){


	total := Sum(20,10);

	if total != 30 {

		t.Errorf("Resultado da soma inválido")
	}
} 