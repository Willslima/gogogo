package contas

import (
	"../clientes"
)

type Contas struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *Contas) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente"
	}
}

func (c *Contas) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "O valor do depósito deve ser maior que zero", c.Saldo
	}
}

func (c *Contas) TransferenciaEntreContas(valorDaTransferencia float64, contaDeDestino *Contas) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
