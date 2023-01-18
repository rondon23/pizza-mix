package models

type Estoque struct {
	Id          int     `json:"id"`
	Descricao   string  `json:"descricao"`
	Quantidade  int     `json:"quantidade"`
	ValorCompra float64 `json:"valor_compra"`
}
