package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("O valor de i: %d hue \r\n", i)
	}

	valor := 0
	teste := true

	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O numero: ", valor)
	}

	for {
		valor--
		fmt.Println("O numero: ", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Eu adoro escrever programas usando Go"
	for indice, letra := range texto {
		fmt.Printf("Texto [%d] = %q\r\n", indice, letra)
	}

}
