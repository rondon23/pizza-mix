package models

import (
	"database/sql"
)

// Model is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Produto is the type for pizza
type Produto struct {
	ID           int     `json: "id"`
	NomeProduto  string  `json:"nome_produto"`
	Descricao    string  `json:"descricao"`
	ValorProduto float32 `json:"valor_produto"`
	FotoProduto  string  `json:"foto_produto"`
}

// Funcionario is the type for pizza
type Funcionario struct {
	ID     int    `json: "id"`
	Nome   string `json: "nome"`
	Rg     string `json: "rg"`
	Cpf    string `json: "cpf"`
	Rua    string `json: "rua"`
	Numero int    `json: "numero"`
	Bairro string `json: "bairro"`
	Cargo  string `json: "cargo"`
}

type Estoque struct {
	ID          int     `json: "id"`
	Descricao   string  `json:"descricao"`
	Quantidade  int     `json:"quantidade"`
	ValorCompra float64 `json:"valor_compra"`
}

type Fornecedor struct {
	ID       int    `json: "id"`
	Nome     string `json: "nome"`
	Endereco string `json: "endereco"`
	Numero   int    `json:"numero"`
	Bairro   string `json: "bairro"`
	Cnpj     string `json: "cnpj"`
}
