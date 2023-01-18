package models

type Cliente struct {
	Id              int    `json:"cliente_id"`
	NomeCliente     string `json:"nome"`
	Cpf             string `json:"cpf"`
	EnderecoCliente string `json:"endereco"`
	Telefone        string `json:"telefone"`
	Numero          string `json:"numero"`
	Bairro          string `json:"bairro"`
	Email           string `json:"email"`
}
