package model

// Imovel struct
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

// SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor devolve o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
