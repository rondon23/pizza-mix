package domain

import (
	"pizza-backend/errs"
	"pizza-backend/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type CarrinhoRepositoryDb struct {
	client *sqlx.DB
}

func NewCarrinhoRepositoryDb(dbClient *sqlx.DB) CarrinhoRepositoryDb {
	return CarrinhoRepositoryDb{dbClient}
}

func (c CarrinhoRepositoryDb) ById(carrinho_id string) (*Carrinho, *errs.AppError) {
	sqlGetCarrinho := "SELECT carrinho_id, produto_id, preco, quantidade, total, nome_produto, sub_total FROM carrinho WHERE carrinho_id = ?"

	var carrinho Carrinho
	err := c.client.Get(&carrinho, sqlGetCarrinho, carrinho_id)
	if err != nil {
		logger.Error("Error while getting carrinho information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &carrinho, nil
}

func (d CarrinhoRepositoryDb) Save(p Carrinho) (*Carrinho, *errs.AppError) {
	sqlInsert := "INSERT INTO carrinho (produto_id, preco, quantidade, total, nome_produto, sub_total) VALUES(?, ?, ?, ?, ?, ?);"

	result, err := d.client.Exec(sqlInsert, p.CodProduto, p.Preco, p.Quantidade, p.Total, p.NomeProduto, p.SubTotal)
	if err != nil {
		logger.Error("Error while creating new carrinho: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new carinho: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	p.Id = strconv.FormatInt(id, 10)
	return &p, nil
}
