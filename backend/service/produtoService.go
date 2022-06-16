package service

import (
	"pizza-backend/domain"
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type ProdutoService interface {
	GetProduto(string) (*dto.NewProdutoResponse, *errs.AppError)
	GetAllProduto() ([]dto.NewProdutoResponse, *errs.AppError)
	NewProduto(dto.NewProdutoRequest) (*dto.NewProdutoResponse, *errs.AppError)
	UpdateProduto(req dto.NewProdutoResponse) (int64, *errs.AppError)
	DeleteProduto(string) (int64, *errs.AppError)
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

func (s DefaultProdutoService) GetAllProduto() ([]dto.NewProdutoResponse, *errs.AppError) {
	produtos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.NewProdutoResponse, 0)

	for _, c := range produtos {
		response = append(response, c.ToDto())
	}

	return response, nil
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

func (s DefaultProdutoService) UpdateProduto(req dto.NewProdutoResponse) (int64, *errs.AppError) {
	produto := domain.NewProdutoUpdate(req.Id, req.NomeProduto, req.Descricao, req.ValorProduto, req.FotoProduto)
	if newProduto, err := s.repo.Update(produto); err != nil {
		return 0, err
	} else {
		return newProduto, nil
	}
}

func (s DefaultProdutoService) DeleteProduto(id string) (int64, *errs.AppError) {
	p, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}
	return p, nil
}

func NewProdutoService(repository domain.ProdutoRepository) DefaultProdutoService {
	return DefaultProdutoService{repository}
}
