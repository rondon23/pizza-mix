package domain

import (
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type Produto struct {
	Id           string  `db:"produto_id"`
	NomeProduto  string  `db:"nome_produto"`
	Descricao    string  `db:"descricao"`
	ValorProduto float32 `db:"valor_produto"`
	FotoProduto  string  `db:"foto_produto"`
}

func (p Produto) ToNewProdutoResponseDto() *dto.NewProdutoResponse {
	return &dto.NewProdutoResponse{p.Id}
}

type ProdutoRepository interface {
	ById(string) (*Produto, *errs.AppError)
	Save(p Produto) (*Produto, *errs.AppError)
}

func (p Produto) ToDto() dto.NewProdutoResponse {
	return dto.NewProdutoResponse{
		Id:           p.Id,
		NomeProduto:  p.NomeProduto,
		Descricao:    p.Descricao,
		ValorProduto: p.ValorProduto,
		FotoProduto:  p.FotoProduto,
	}
}

func NewProduto(nomeProduto, descricao string, valorProduto float32, fotoProduto string) Produto {
	return Produto{
		NomeProduto:  nomeProduto,
		Descricao:    descricao,
		ValorProduto: valorProduto,
		FotoProduto:  fotoProduto,
	}
}
