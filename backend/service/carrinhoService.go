package service

import (
	"pizza-backend/domain"
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type CarrinhoService interface {
	GetCarrinho(id string) (*dto.NewCarrinhoResponse, *errs.AppError)
	NewCarrinho(req dto.NewCarrinhoRequest) (*dto.NewCarrinhoResponse, *errs.AppError)
	GetAllCarrinho() ([]dto.NewCarrinhoResponse, *errs.AppError)
}

type DefaultCarrinhoService struct {
	repo domain.CarrinhoRepository
}

func (c DefaultCarrinhoService) GetCarrinho(id string) (*dto.NewCarrinhoResponse, *errs.AppError) {
	p, err := c.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := p.ToDto()
	return &response, nil
}

func (s DefaultCarrinhoService) GetAllCarrinho() ([]dto.NewCarrinhoResponse, *errs.AppError) {
	carrinhos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.NewCarrinhoResponse, 0)

	for _, c := range carrinhos {
		response = append(response, c.ToDto())
	}

	return response, nil
}

func (s DefaultCarrinhoService) NewCarrinho(req dto.NewCarrinhoRequest) (*dto.NewCarrinhoResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	carrinho := domain.NewCarrinho(req.CodProduto, req.Quantidade, req.Preco, req.SubTotal, req.Total, req.NomeProduto)
	if newCarrinho, err := s.repo.Save(carrinho); err != nil {
		return nil, err
	} else {
		return newCarrinho.ToNewCarrinhoResponseDto(), nil
	}
}

func NewCarrinhoService(repository domain.CarrinhoRepository) DefaultCarrinhoService {
	return DefaultCarrinhoService{repository}
}
