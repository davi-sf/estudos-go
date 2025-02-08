package main

import (
	"fmt"
)

func main() {

	fmt.Print("Insira um numero para verificar se eh primo: \n")
	var numero int
	fmt.Scan(&numero)

	if numero <= 1 {
		fmt.Print("nao eh primo")
	}

	var primo bool = true

	for i := 2; i <= numero/2; i++ {
		if numero%i == 0 {
			primo = false
			break
		}
	}

	if primo {
		fmt.Print("eh primo")
	} else {
		fmt.Print("nao eh primo")
	}

}
