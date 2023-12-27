package contas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockContaCorrente = ContaCorrente{}
var mockContaCorrenteTransferencia = ContaCorrente{}

func TestSacar(t *testing.T) {
	mockContaCorrente.Depositar(100) 
	result := mockContaCorrente.Sacar(50)

	assert.Equal(t, "Saldo realizado com sucesso", result)
}

func TestSacarSaldoInvalid(t *testing.T) {
	mockContaCorrente.Depositar(100) 
	result := mockContaCorrente.Sacar(200)

	assert.Equal(t, "Saldo insuficiente", result)
}

func TestDepositar(t *testing.T) {
	status, saldo := mockContaCorrente.Depositar(10)

	assert.Equal(t, "Deposito realizado com sucesso. Saldo = ", status)
	assert.Equal(t, 10.0, saldo)
}

func TestDepositarInvalid(t *testing.T) {
	status, saldo := mockContaCorrente.Depositar(-10)

	assert.Equal(t, "Valor do deposito menor que zero. Saldo = ", status)
	assert.Equal(t, 0.0, saldo)
}

func TestTranferencia(t *testing.T){
	mockContaCorrente.Depositar(100) 
	status := mockContaCorrente.Transferir(10.0, &mockContaCorrenteTransferencia)
	
	assert.True(t, status)
}

func TestTranferenciaInvalid(t *testing.T){
	status := mockContaCorrente.Transferir(10.0, &mockContaCorrenteTransferencia)
	
	assert.False(t, status)
}

func TestObterSaldo(t *testing.T) {
	saldo := mockContaCorrente.ObterSaldo()

	assert.Equal(t, 0.0, saldo)
}
