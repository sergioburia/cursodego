package main

import (
	"fmt"
)

func main() {
	meses := 1
	situacao := false
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	} else {
		fmt.Println("Ele está adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive na cidade Teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situacao do cliente? ", descricao)
		return
	}
	fmt.Println("Obrigado!")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "Ele esta devendo"
		return
	}
	descricao = "Ele não esta devendo"
	return
}
