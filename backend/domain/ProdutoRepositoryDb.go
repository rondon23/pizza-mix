package domain

import "github.com/jmoiron/sqlx"

type ProdutoRepositoryDb struct {
	client *sqlx.DB
}
