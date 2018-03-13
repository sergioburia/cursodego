package main

import (
	"fmt"
)

var (
	//Endereco é um valor publico
	Endereco string
	//Telefone é um valor publico
	Telefone           string
	quantiade, estoque int     //quantidade = 0
	comprou            bool    //comprou = false
	valor              float64 //valor = 0.00
	palavras           rune
)

func main() {
	teste := "valor do teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantiade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("PALAVRAS: %v\r\n", palavras)
	fmt.Printf("Teste: %v\r\n", teste)
}
