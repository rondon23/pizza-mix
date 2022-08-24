package domain

import (
	"pizza-backend/dto"
	"pizza-backend/errs"
)

type Cliente struct {
	Id              int    `db:"cliente_id"`
	NomeCliente     string `db:"nome"`
	Cpf             string `db:"cpf"`
	EnderecoCliente string `db:"endereco"`
	Telefone        string `db:"telefone"`
	Numero          string `db:"numero"`
	Bairro          string `db:"bairro"`
	Email           string `db:"email"`
}

func (c Cliente) ToNewClienteResponseDto() *dto.NewClienteResponse {
	return &dto.NewClienteResponse{c.Id, c.NomeCliente, c.Cpf, c.EnderecoCliente, c.Telefone, c.Numero, c.Bairro, c.Email}
}

type ClienteRepository interface {
	GetAll() ([]Cliente, *errs.AppError)
	Save(c Cliente) (*Cliente, *errs.AppError)
}

func NewCliente(nomeCliente, cpf, enderecoCliente, telefone, numero, bairro, email string) Cliente {
	return Cliente{
		NomeCliente:     nomeCliente,
		Cpf:             cpf,
		EnderecoCliente: enderecoCliente,
		Telefone:        telefone,
		Numero:          numero,
		Bairro:          bairro,
		Email:           email,
	}
}

func (c Cliente) ToDto() dto.NewClienteResponse {
	return dto.NewClienteResponse{
		Id:              c.Id,
		NomeCliente:     c.NomeCliente,
		Cpf:             c.Cpf,
		EnderecoCliente: c.EnderecoCliente,
		Telefone:        c.Telefone,
		Numero:          c.Numero,
		Bairro:          c.Bairro,
		Email:           c.Email,
	}
}
