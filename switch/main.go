package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 5
	fmt.Print("O numero ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("tres")
	}

	fmt.Println("Voce é da familia do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println(runtime.GOOS)
	default:
		fmt.Println("Nope")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("OK Voce pode dormir")
	default:
		fmt.Println("Dia de trabalho")
	}

	numero = 1
	fmt.Println("Este numero cabe num digito?")
	switch {
	case numero < 10:
		fmt.Println("SIM")
	case numero >= 10 && numero < 100:
		fmt.Println("2 digitos")
	case numero >= 100:
		fmt.Println("maior que 100")
	}
}
