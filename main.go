package golang

import (
	"github.com/yanvic/aaaaa/conta"
)

type ContaConrrente struct {
	Titular     conta.Titular
	NumeroConta int
	Saldo       float64
}

func (c *ContaConrrente) Saque(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "saldo realizado"
	} else {
		return "saldo insuficiente"
	}

}

func (c *ContaConrrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "deposito realizado", c.Saldo
	} else {
		return "o valor do deposito é menor que zero", c.Saldo
	}
}

func (c *ContaConrrente) Tranferir(valorTransferencia float64, contaDestino *ContaConrrente) bool {
	if valorTransferencia < c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
