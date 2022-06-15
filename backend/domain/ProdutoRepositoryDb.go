package domain

import (
	"pizza-backend/errs"

	"pizza-backend/logger"

	"github.com/jmoiron/sqlx"
)

type ProdutoRepositoryDb struct {
	client *sqlx.DB
}

func (p ProdutoRepositoryDb) ById(produto_id string) (*Produto, *errs.AppError) {
	sqlGetProduto := "SELECT produto_id, nome_produto, descricao, valor_produto, foto_produto FROM produto where produto_id = ?"
	var produto Produto
	err := p.client.Get(&produto, sqlGetProduto, produto_id)
	if err != nil {
		logger.Error("Error while fetching product information: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &produto, nil
}

func NewProdutoRepositoryDb(dbClient *sqlx.DB) ProdutoRepositoryDb {
	return ProdutoRepositoryDb{dbClient}
}
