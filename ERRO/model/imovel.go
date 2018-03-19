package model

import (
	"errors"
)

// Imovel struct
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//ErrValorDeImovelInvalido muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor do imovel é invalido")

//ErrValordeImovelMuitoAlto muito alto
var ErrValordeImovelMuitoAlto = errors.New("VAlor muito alto")

// SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 999999 {
		err = ErrValordeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor devolve o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
