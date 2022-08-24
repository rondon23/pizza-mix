package service

import (
	"pizza-backend/domain"
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type ClienteService interface {
	GetAllCliente() ([]dto.NewClienteResponse, *errs.AppError)
	NewCliente(req dto.NewClienteRequest) (*dto.NewClienteResponse, *errs.AppError)
}

type DefaultClienteService struct {
	repo domain.ClienteRepository
}

func (s DefaultClienteService) GetAllCliente() ([]dto.NewClienteResponse, *errs.AppError) {
	clientes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.NewClienteResponse, 0)

	for _, c := range clientes {
		response = append(response, c.ToDto())
	}

	return response, nil
}

func (s DefaultClienteService) NewCliente(req dto.NewClienteRequest) (*dto.NewClienteResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	cliente := domain.NewCliente(req.NomeCliente, req.Cpf, req.EnderecoCliente, req.Telefone, req.Numero, req.Bairro, req.Email)
	if newCliente, err := s.repo.Save(cliente); err != nil {
		return nil, err
	} else {
		return newCliente.ToNewClienteResponseDto(), nil
	}
}

func NewClienteService(repository domain.ClienteRepository) DefaultClienteService {
	return DefaultClienteService{repository}
}
