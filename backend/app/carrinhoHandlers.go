package app

import (
	"encoding/json"
	"net/http"
	"pizza-backend/dto"
	"pizza-backend/service"

	"github.com/gorilla/mux"
)

type CarrinhoHandlers struct {
	service service.CarrinhoService
}

func (ch *CarrinhoHandlers) GetCarrinho(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["carrinho_id"]

	carrinho, err := ch.service.GetCarrinho(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, carrinho)
	}
}

func (ch CarrinhoHandlers) NewCarrinho(w http.ResponseWriter, r *http.Request) {
	var request dto.NewCarrinhoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		produto, appError := ch.service.NewCarrinho(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, produto)
		}
	}
}
