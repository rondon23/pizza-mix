package models

import "time"

type Compras struct {
	Id            int       `json:"id"`
	DataCompra    time.Time `json:"data_compra"`
	Fornecedor    string    `json:"fornecedor"`
	CodProduto    int       `json:"cod_produto"`
	PrecoUnitario float64   `json:"preco_unitario"`
	ValorCompra   float64   `json:"valor_compra"`
}
