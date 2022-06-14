package domain

type Estoque struct {
	ID          int     `json:"id"`
	Descricao   string  `json:"descricao"`
	Quantidade  int     `json:"quantidade"`
	ValorCompra float64 `json:"valor_compra"`
}
