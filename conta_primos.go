package main

import "fmt"

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	for i := 2; i <= numero/2; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Print("insira um limite para verificar a quantidade de numeros primos de 1 ate esse valor: \n ")

	var limite int
	fmt.Scan(&limite)
	var contador int

	for i := 1; i <= limite; i++ {
		if ehPrimo(i) {
			contador++
		}
	}

	fmt.Print(contador)

}
