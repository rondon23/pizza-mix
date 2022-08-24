package dto

type NewClienteResponse struct {
	Id              int    `json:"cliente_id"`
	NomeCliente     string `json:"nome_cliente"`
	Cpf             string `json:"cpf"`
	EnderecoCliente string `json:"endereco_cliente"`
	Telefone        string `json:"telefone"`
	Numero          string `json:"numero"`
	Bairro          string `json:"bairro"`
	Email           string `json:"email"`
}
