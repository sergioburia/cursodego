package matematica

//Calculo executa qualqeur tipo de calculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador publico
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor auhehauea
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto auhehasasasaue aeaiehauehia
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
