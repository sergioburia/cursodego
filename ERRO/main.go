package main

import (
	"encoding/json"
	"fmt"

	"github.com/sergioburia/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25

	if err := casa.SetValor(1); err != nil {
		fmt.Println("[main] Erro no valor da casa: ", err)
		if err == model.ErrValordeImovelMuitoAlto {
			fmt.Println("Corretor, verifique avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O Valor da casa é: %d\r\n", casa.GetValor())
	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Erro no JSON: ", err.Error)
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))
}
