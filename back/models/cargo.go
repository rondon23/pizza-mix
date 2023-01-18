package models

type Cargo struct {
	Id         uint         `json:"id"`
	Nome       string       `json:"nome"`
	Permissoes []Permissoes `json:"permissoes" gorm:"many2many:cargo_permissoes"`
}
