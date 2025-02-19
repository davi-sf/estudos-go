package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Operacao interface {
	Executar(saldo float64) (float64, error)
}

type Deposito struct {
	Valor float64
}

func (deposito Deposito) Executar(saldo float64) (float64, error) {
	if deposito.Valor <= 0 {
		return saldo, fmt.Errorf("valor invalido, tente novamente")
	}
	return saldo + deposito.Valor, nil
}

type Saque struct {
	Valor float64
}

func (saque Saque) Executar(saldo float64) (float64, error) {
	if saque.Valor > saldo {
		return saldo, fmt.Errorf("saldo insuficiente")
	}
	return saldo - saque.Valor, nil
}

type Banco struct {
	Saldo float64
}

func (banco *Banco) RealizarOperacao(op Operacao) {
	novoSaldo, err := op.Executar(banco.Saldo)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	banco.Saldo = novoSaldo
	writeFloatToFile(banco.Saldo, "saldo.txt")
	fmt.Println("Operação realizada com sucesso!")
}

func getFloatToFile(fileName string, defaultValueReturn float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValueReturn, errors.New("File not found")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {

	saldoInicial, _ := getFloatToFile("saldo.txt", 0.0)
	banco := Banco{Saldo: saldoInicial}

	for {

		menu()
		var opt int
		fmt.Scan(&opt)

		switch opt {
		case 1:
			fmt.Printf("O saldo disponível eh: %.2f\n", banco.Saldo)
		case 2:
			var valorDepositado float64
			fmt.Print("Insira o valor a ser depositado: ")
			fmt.Scan(&valorDepositado)
			banco.RealizarOperacao(Deposito{Valor: valorDepositado})
		case 3:
			var valorSaque float64
			fmt.Print("Insira o valor a ser sacado: ")
			fmt.Scan(&valorSaque)
			banco.RealizarOperacao(Saque{Valor: valorSaque})
		case 4:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção invalida! Tente novamente.")
		}
	}
}
