package domain

type Cliente struct {
	ID              int    `json:"id"`
	NomeCliente     string `json:"nome_cliente"`
	Cpf             string `json:"cpf"`
	EnderecoCliente string `json:"endereco_cliente"`
	Telefone        string `json:"telefone"`
	Numero          string `json:"numero"`
	Bairro          string `json:"bairro"`
	Email           string `json:"email"`
}
