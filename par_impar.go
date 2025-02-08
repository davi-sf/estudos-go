package main

import "fmt"

func main() {

	fmt.Print("---- Verifica se um numero eh par ou impar ---- \n")

	var n1 int
	fmt.Print("Insira o numero \n")
	fmt.Scan(&n1)
	var msg string
	if n1%2 == 0 {
		msg = "eh par"
	} else {
		msg = "eh impar"
	}

	fmt.Println(msg)

}
