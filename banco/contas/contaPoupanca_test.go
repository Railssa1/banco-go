package contas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockContaPoupanca = ContaPoupanca{}

func TestSacarPoupanca(t *testing.T) {
	mockContaPoupanca.Depositar(100)
	result := mockContaPoupanca.Sacar(50)

	assert.Equal(t, "Saldo realizado com sucesso", result)
}

func TestSacarSaldoInvalidPoupanca(t *testing.T) {
	mockContaPoupanca.Depositar(100)
	result := mockContaPoupanca.Sacar(200)

	assert.Equal(t, "Saldo insuficiente", result)
}

func TestDepositarPoupanca(t *testing.T) {
	status, saldo := mockContaPoupanca.Depositar(10)

	assert.Equal(t, "Deposito realizado com sucesso. Saldo = ", status)
	assert.Equal(t, 10.0, saldo)
}

func TestObterSaldoPoupanca(t *testing.T) {
	saldo := mockContaPoupanca.ObterSaldo()

	assert.Equal(t, 0.0, saldo)
}
