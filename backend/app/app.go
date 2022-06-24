package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pizza-backend/domain"
	"pizza-backend/service"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enviroment variable not defined...")
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	// wiring
	dbClient := getDbClient()
	produtoRepositoryDb := domain.NewProdutoRepositoryDb(dbClient)
	ph := ProdutoHandlers{service: service.NewProdutoService(produtoRepositoryDb)}

	carrinhoRepositoryDb := domain.NewCarrinhoRepositoryDb(dbClient)
	ca := CarrinhoHandlers{service: service.NewCarrinhoService(carrinhoRepositoryDb)}

	// define routes for Produto
	router.
		HandleFunc("/produto/{produto_id:[0-9]+}", ph.GetProduto).
		Methods(http.MethodGet).
		Name("GetProdutoById")

	router.
		HandleFunc("/produto/", ph.GetAllProduto).
		Methods(http.MethodGet).
		Name("GetAllProduto")

	router.
		HandleFunc("/produto/", ph.NewProduto).
		Methods(http.MethodPost).
		Name("NewProduto")

	router.
		HandleFunc("/produto/", ph.UpdateProduto).
		Methods(http.MethodPut).
		Name("UpdateProduto")

	router.
		HandleFunc("/produto/{produto_id:[0-9]+}", ph.DeleteProduto).
		Methods(http.MethodDelete).
		Name("DeleteProduto")

	// define routes for Carrinho
	router.
		HandleFunc("/carrinho/{carrinho_id:[0-9]+}", ca.GetCarrinho).
		Methods(http.MethodGet).
		Name("GetCarrinhoById")

	router.
		HandleFunc("/carrinho/", ca.GetAllCarrinho).
		Methods(http.MethodGet).
		Name("GetAllCarrinho")

	router.
		HandleFunc("/carrinho/", ca.NewCarrinho).
		Methods(http.MethodPost).
		Name("NewCarrinho")

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
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

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
