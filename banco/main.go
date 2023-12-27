package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificaConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificaConta interface {
	Sacar(valorDoSaque float64) string
}

func main() {
	contaPoupanca := contas.ContaPoupanca{}
	contaPoupanca.Depositar(150)
	PagarBoleto(&contaPoupanca, 100)
	fmt.Println(contaPoupanca)

	contaCorrente := contas.ContaCorrente{}
	contaCorrente.Depositar(250)
	PagarBoleto(&contaCorrente, 122)
	fmt.Println(contaCorrente)
}
