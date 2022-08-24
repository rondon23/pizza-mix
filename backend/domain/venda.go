package domain

import "time"

type Venda struct {
	ID          int       `db:"id"`
	DataVenda   time.Time `db:"date_venda"`
	Valor       float32   `db:"valor"`
	TotalVenda  float32   `db:"total_venda"`
	CodCliente  int       `db:"cod_cliente"`
	NomeCliente string    `db:"nome_cliente"`
}
