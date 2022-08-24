package domain

import "time"

type ContasPagar struct {
	ID            int       `db:"id"`
	DataPagamento time.Time `db:"data_pagamento"`
	Valor         float64   `db:"valor"`
}
