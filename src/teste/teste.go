package main

import (
	"fmt"
	"os"
)

func main() {
	// Define uma nova variável de ambiente chamada "MINHA_VARIAVEL" com o valor "meu_valor"
	// os.Setenv("MINHA_VARIAVEL", "valor da variavel")

	// Acessa a variável de ambiente "MINHA_VARIAVEL"
	minhaVariavel := os.Getenv("%WILL%")
	fmt.Println("O valor da variável de ambiente MINHA_VARIAVEL é:", minhaVariavel)
}
