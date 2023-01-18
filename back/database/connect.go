package database

import (
	"github.com/rondon23/pizza-mix/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("developer:1234567@/pizza_mix"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(
		&models.Carrinho{},
		&models.Cliente{},
		&models.Compras{},
		&models.Estoque{},
		&models.Fornecedor{},
		&models.Funcionario{},
		&models.Produto{},
		&models.Venda{},
	)
}
