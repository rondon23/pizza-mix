package models

import "golang.org/x/crypto/bcrypt"

type Funcionario struct {
	Id      int    `json:"id"`
	Nome    string `json:"nome"`
	Rg      string `json:"rg"`
	Cpf     string `json:"cpf"`
	Rua     string `json:"rua"`
	Numero  int    `json:"numero"`
	Bairro  string `json:"bairro"`
	Email   string `json:"email" gorm:"unique"`
	Senha   []byte `json:"-"`
	CargoId uint   `json:"cargo_id"`
	Cargo   Cargo  `json:"cargo" gorm:"foreignKey:CargoId"`
}

func (funcionario *Funcionario) DefinirSenha(senha string) {
	hashedSenha, _ := bcrypt.GenerateFromPassword([]byte(senha), 14)
	funcionario.Senha = hashedSenha
}

func (funcionario *Funcionario) CompararSenhas(senha string) error {
	return bcrypt.CompareHashAndPassword(funcionario.Senha, []byte(senha))
}
