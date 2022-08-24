package domain

import (
	"pizza-backend/errs"
	"pizza-backend/logger"

	"github.com/jmoiron/sqlx"
)

type ClienteRepositoryDb struct {
	client *sqlx.DB
}

func NewClienteRepositoryDb(dbClient *sqlx.DB) ClienteRepositoryDb {
	return ClienteRepositoryDb{dbClient}
}

func (c ClienteRepositoryDb) GetAll() ([]Cliente, *errs.AppError) {
	sqlGetCliente := "SELECT cliente_id, nome, cpf, endereco, telefone, numero, bairro, email FROM cliente"

	clientes := make([]Cliente, 0)

	err := c.client.Select(&clientes, sqlGetCliente)
	if err != nil {
		logger.Error("Error while fetching carrinho information: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return clientes, nil
}

func (d ClienteRepositoryDb) Save(c Cliente) (*Cliente, *errs.AppError) {
	sqlInsert := "INSERT INTO pizza_mix.cliente (nome, cpf, endereco, telefone, numero, bairro, email) VALUES(?, ?, ?, ?, ?, ?, ?);"

	result, err := d.client.Exec(sqlInsert, c.NomeCliente, c.Cpf, c.EnderecoCliente, c.Telefone, c.Numero, c.Bairro, c.Email)
	if err != nil {
		logger.Error("Error while creating new cliente: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new cliente: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	c.Id = int(id)
	return &c, nil
}
