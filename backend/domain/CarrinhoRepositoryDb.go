package domain

import "github.com/jmoiron/sqlx"

type CarrinhoRepositoryDb struct {
	client *sqlx.DB
}
