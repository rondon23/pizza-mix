package domain

import (
	"pizza-backend/errs"
	"strconv"

	"pizza-backend/logger"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
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

func (d ProdutoRepositoryDb) Save(p Produto) (*Produto, *errs.AppError) {
	sqlInsert := "INSERT INTO pizza_mix.produto (nome_produto, descricao, valor_produto, foto_produto) VALUES(?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, p.NomeProduto, p.Descricao, p.ValorProduto, p.FotoProduto)
	if err != nil {
		logger.Error("Error while creating new product: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new product: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	p.Id = strconv.FormatInt(id, 10)
	return &p, nil
}
