package main

import (
	"encoding/json"
	"fmt"

	"github.com/sergioburia/structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)
	//casa.valor = 2

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O Valor da casa é: %+d\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)

	fmt.Println("A casa em JSON: ", string(objJSON))

	ii := model.Imovel{X: 2, Y: 1, Nome: "batatinha"}
	fmt.Printf("batata %d\r\n", ii.GetValor())

}
