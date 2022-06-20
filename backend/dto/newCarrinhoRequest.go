package dto

import "pizza-backend/errs"

type NewCarrinhoRequest struct {
	CodProduto  int     `json:"produto_id"`
	Preco       float32 `json:"preco"`
	Quantidade  int     `json:"quantidade"`
	Total       float32 `json:"total"`
	NomeProduto string  `json:"nome_produto"`
	SubTotal    float32 `json:"sub_total"`
}

func (r NewCarrinhoRequest) Validate() *errs.AppError {
	return nil
}
