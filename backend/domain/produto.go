package domain

import (
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type Produto struct {
	ID           int     `json:"produto_id"`
	NomeProduto  string  `json:"nome_produto"`
	Descricao    string  `json:"descricao"`
	ValorProduto float32 `json:"valor_produto"`
	FotoProduto  string  `json:"foto_produto"`
}

type ProdutoRepository interface {
	ById(string) (*Produto, *errs.AppError)
}

func (p Produto) ToDto() dto.ProdutoResponse {
	return dto.ProdutoResponse{
		Id:           p.ID,
		NomeProduto:  p.NomeProduto,
		Descricao:    p.Descricao,
		ValorProduto: p.ValorProduto,
		FotoProduto:  p.FotoProduto,
	}
}
