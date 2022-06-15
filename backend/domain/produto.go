package domain

import (
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type Produto struct {
	Id           int     `db:"produto_id"`
	NomeProduto  string  `db:"nome_produto"`
	Descricao    string  `db:"descricao"`
	ValorProduto float32 `db:"valor_produto"`
	FotoProduto  string  `db:"foto_produto"`
}

type ProdutoRepository interface {
	ById(string) (*Produto, *errs.AppError)
}

func (p Produto) ToDto() dto.ProdutoResponse {
	return dto.ProdutoResponse{
		Id:           p.Id,
		NomeProduto:  p.NomeProduto,
		Descricao:    p.Descricao,
		ValorProduto: p.ValorProduto,
		FotoProduto:  p.FotoProduto,
	}
}
