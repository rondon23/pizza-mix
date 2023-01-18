package models

type Fornecedor struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Numero   int    `json:"numero"`
	Bairro   string `json:"bairro"`
	Cnpj     string `json:"cnpj"`
}
