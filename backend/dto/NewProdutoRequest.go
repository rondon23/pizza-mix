package dto

import (
	"pizza-backend/errs"
)

type NewProdutoRequest struct {
	NomeProduto  string  `json:"nome_produto"`
	Descricao    string  `json:"descricao"`
	ValorProduto float32 `json:"valor_produto"`
	FotoProduto  string  `json:"foto_produto"`
}

func (r NewProdutoRequest) Validate() *errs.AppError {
	return nil
}
