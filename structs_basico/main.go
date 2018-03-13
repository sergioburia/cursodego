package main

import (
	"fmt"
)

//Imovel é uma struct
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %v\r\n", casa)

	apartamento := Imovel{17, 56, "meu ap", 760000}
	fmt.Printf("O apartamento é %v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     22,
	}
	fmt.Printf("A chacara é %v\r\n", chacara)

	casa.Nome = "Lar doe Lar"
	casa.valor = 444444
	casa.X = 10
	casa.Y = 12
	fmt.Printf("A casa é: %v\r\n", casa)
}
