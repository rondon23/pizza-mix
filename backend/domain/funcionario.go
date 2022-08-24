package domain

type Funcionario struct {
	ID     int    `db:"id"`
	Nome   string `db:"nome"`
	Rg     string `db:"rg"`
	Cpf    string `db:"cpf"`
	Rua    string `db:"rua"`
	Numero int    `db:"numero"`
	Bairro string `db:"bairro"`
	Cargo  string `db:"cargo"`
}
