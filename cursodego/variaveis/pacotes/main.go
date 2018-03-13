package main

import (
	"fmt"

	"github.com/sergioburia/cursodego/variaveis/pacotes/operadora"
	"github.com/sergioburia/cursodego/variaveis/pacotes/prefixo"
)

//NomeDoUsuario do sistema
var NomeDoUsuario = "Sergio"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo: %d\r\n", prefixo.Capital)
	fmt.Printf("Operadora: %d\r\n", operadora.NomeOperadora)
}
