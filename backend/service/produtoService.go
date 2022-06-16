package service

import (
	"pizza-backend/domain"
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type ProdutoService interface {
	GetProduto(string) (*dto.NewProdutoResponse, *errs.AppError)
	NewProduto(dto.NewProdutoRequest) (*dto.NewProdutoResponse, *errs.AppError)
}

type DefaultProdutoService struct {
	repo domain.ProdutoRepository
}

func (s DefaultProdutoService) GetProduto(id string) (*dto.NewProdutoResponse, *errs.AppError) {
	p, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := p.ToDto()
	return &response, nil
}

func (s DefaultProdutoService) NewProduto(req dto.NewProdutoRequest) (*dto.NewProdutoResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	produto := domain.NewProduto(req.NomeProduto, req.Descricao, req.ValorProduto, req.FotoProduto)
	if newProduto, err := s.repo.Save(produto); err != nil {
		return nil, err
	} else {
		return newProduto.ToNewProdutoResponseDto(), nil
	}
}

func NewProdutoService(repository domain.ProdutoRepository) DefaultProdutoService {
	return DefaultProdutoService{repository}
}
