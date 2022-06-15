package service

import (
	"pizza-backend/domain"
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type ProdutoService interface {
	GetProduto(string) (*dto.ProdutoResponse, *errs.AppError)
}

type DefaultProdutoService struct {
	repo domain.ProdutoRepository
}

func (s DefaultProdutoService) GetProduto(id string) (*dto.ProdutoResponse, *errs.AppError) {
	p, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := p.ToDto()
	return &response, nil
}

func NewProdutoService(repository domain.ProdutoRepository) DefaultProdutoService {
	return DefaultProdutoService{repository}
}
