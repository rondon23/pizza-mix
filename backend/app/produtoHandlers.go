package app

import (
	"encoding/json"
	"net/http"
	"pizza-backend/service"

	"github.com/gorilla/mux"
)

type ProdutoHandlers struct {
	service service.ProdutoService
}

func (ch *ProdutoHandlers) getProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	produto, err := ch.service.GetProduto(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, produto)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
