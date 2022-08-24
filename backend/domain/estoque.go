package domain

type Estoque struct {
	ID          int     `db:"id"`
	Descricao   string  `db:"descricao"`
	Quantidade  int     `db:"quantidade"`
	ValorCompra float64 `db:"valor_compra"`
}
