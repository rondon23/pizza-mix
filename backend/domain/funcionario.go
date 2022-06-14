package domain

type Funcionario struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Rg     string `json:"rg"`
	Cpf    string `json:"cpf"`
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Bairro string `json:"bairro"`
	Cargo  string `json:"cargo"`
}
