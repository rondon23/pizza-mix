package dto

import "pizza-backend/errs"

type NewClienteRequest struct {
	Id              string `json:"id_cliente"`
	NomeCliente     string `json:"nome_cliente"`
	Cpf             string `json:"cpf"`
	EnderecoCliente string `json:"endereco_cliente"`
	Telefone        string `json:"telefone"`
	Numero          string `json:"numero"`
	Bairro          string `json:"bairro"`
	Email           string `json:"email"`
}

func (r NewClienteRequest) Validate() *errs.AppError {
	return nil
}
