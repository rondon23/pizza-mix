package domain

import (
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type Carrinho struct {
	Id          string  `db:"carrinho_id"`
	CodProduto  int     `db:"produto_id"`
	Preco       float32 `db:"preco"`
	Quantidade  int     `db:"quantidade"`
	Total       float32 `db:"total"`
	NomeProduto string  `db:"nome_produto"`
	SubTotal    float32 `db:"sub_total"`
}

func (p Carrinho) ToNewCarrinhoResponseDto() *dto.NewCarrinhoResponse {
	return &dto.NewCarrinhoResponse{p.Id, p.CodProduto, p.Preco, p.Quantidade, p.Total, p.NomeProduto, p.SubTotal}
}

type CarrinhoRepository interface {
	ById(string) (*Carrinho, *errs.AppError)
	Save(p Carrinho) (*Carrinho, *errs.AppError)
	GetAll() ([]Carrinho, *errs.AppError)
}

func NewCarrinho(codProduto, quantidade int, preco, subTotal, total float32, nomeProduto string) Carrinho {
	return Carrinho{
		CodProduto:  codProduto,
		Preco:       preco,
		Quantidade:  quantidade,
		Total:       total,
		NomeProduto: nomeProduto,
		SubTotal:    subTotal,
	}
}

func (p Carrinho) ToDto() dto.NewCarrinhoResponse {
	return dto.NewCarrinhoResponse{
		Id:          p.Id,
		CodProduto:  p.CodProduto,
		Preco:       p.Preco,
		Quantidade:  p.Quantidade,
		Total:       p.Total,
		NomeProduto: p.NomeProduto,
		SubTotal:    p.SubTotal,
	}
}
