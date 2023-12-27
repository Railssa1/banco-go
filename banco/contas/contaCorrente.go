package contas

import (
	c "banco/clientes"
)

type ContaCorrente struct {
	Titular                    c.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque <= c.saldo && valorDoSaque > 0 {
		c.saldo -= valorDoSaque
		return "Saldo realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. Saldo = ", c.saldo
	}

	return "Valor do deposito menor que zero. Saldo = ", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia <= c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Sacar(valorTransferencia)
		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
