package main

import (
	"fmt"
	"./contas"
	"./clientes"
)

func main() {
	conta := contas.Contas{Titular: clientes.Titular {
		Nome:      "Will",
		CPF:       "111.222.333-44",
		Profissao: "Desenvolvedor"}, 
		Agencia: 123, Conta: 12345, Saldo: 100}

	fmt.Println(conta)
}
