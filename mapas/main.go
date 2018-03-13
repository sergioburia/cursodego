package main

import (
	"fmt"

	"github.com/sergioburia/structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e achei isso: \r\n%+v\r\n%+v\r\n", achou, imovel)
		//fmt.Printf("Sim, e achei isso: \r\n%+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 22
	apto.Y = 33
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Itens no cache: ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}
	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}
	fmt.Println("Itens no cache: ", len(cache))
}
