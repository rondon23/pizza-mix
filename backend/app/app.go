package app

import (
	"fmt"
	"net/http"
	"os"
	"pizza-backend/domain"
	"pizza-backend/logger"
	"pizza-backend/service"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Enviroment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// wiring
	dbClient := getDbClient()
	produtoRepositoryDb := domain.NewProdutoRepositoryDb(dbClient)
	ph := ProdutoHandlers{service.NewProdutoService(produtoRepositoryDb)}

	// define routes
	router.HandleFunc("/produto/{produto_id:[0-9]+}", ph.getProduto).
		Methods(http.MethodGet).
		Name("GetProdutoById")
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	//See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
