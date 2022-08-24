package domain

type Fornecedor struct {
	ID       int    `db:"id"`
	Nome     string `db:"nome"`
	Endereco string `db:"endereco"`
	Numero   int    `db:"numero"`
	Bairro   string `db:"bairro"`
	Cnpj     string `db:"cnpj"`
}
