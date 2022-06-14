package domain

import "time"

type ContasPagar struct {
	ID            int       `json:"id"`
	DataPagamento time.Time `json:"data_pagamento"`
	Valor         float64   `json:"valor"`
}
