package domain

import "time"

type Venda struct {
	ID          int       `json:"id"`
	DataVenda   time.Time `json:"date_venda"`
	Valor       float32   `json:"valor"`
	TotalVenda  float32   `json:"total_venda"`
	CodCliente  int       `json:"cod_cliente"`
	NomeCliente string    `json:"nome_cliente"`
}
