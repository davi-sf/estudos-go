package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var msg string
	var contador int = 0

	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("frase ou palavra para verificar a quantidade de vogais \n")
	sc.Scan()
	msg = sc.Text()

	strings.ToLower(msg)
	for i := 0; i < len(msg); i++ {
		if msg[i] == 'a' || msg[i] == 'e' || msg[i] == 'i' || msg[i] == 'o' || msg[i] == 'u' {
			contador++
		}
	}

	println("a quantidade de vogais eh: ", contador)

}
