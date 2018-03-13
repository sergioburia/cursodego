package main

import (
	"fmt"

	"github.com/sergioburia/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 3, 4)
	fmt.Printf("Resultado: %d\r\n", resultado)
	resultado = matematica.Soma(3, 2)
	fmt.Printf("Resultado: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 4)
	fmt.Printf("Resultado: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(11, 4)
	fmt.Printf("Resultado: %d\r\n", resultado)
	fmt.Printf("Resto: %d\r\n", resto)

}
