package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "O valor do depósito deve ser maior que zero", c.saldo
	}
}

func main() {
	conta := ContaCorrente{titular: "Willian", agencia: 123, conta: 12345, saldo: 125.5}

	status, valor := conta.Depositar(51.5)
	fmt.Println("Status:", status)
	fmt.Println("Valor:", valor)
}
