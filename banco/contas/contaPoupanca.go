package contas

import c "banco/clientes"

type ContaPoupanca struct {
	Titular                              c.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	if valorDoSaque <= c.saldo && valorDoSaque > 0 {
		c.saldo -= valorDoSaque
		return "Saldo realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. Saldo = ", c.saldo
	}

	return "Valor do deposito menor que zero. Saldo = ", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

