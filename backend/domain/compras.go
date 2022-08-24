package domain

import "time"

type Compras struct {
	ID            int       `db:"id"`
	DataCompra    time.Time `db:"data_compra"`
	Fornecedor    string    `db:"fornecedor"`
	CodProduto    int       `db:"cod_produto"`
	PrecoUnitario float64   `db:"preco_unitario"`
	ValorCompra   float64   `db:"valor_compra"`
}
