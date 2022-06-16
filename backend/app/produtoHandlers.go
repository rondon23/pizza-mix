package app

import (
	"encoding/json"
	"net/http"
	"pizza-backend/dto"
	"pizza-backend/service"

	"github.com/gorilla/mux"
)

type ProdutoHandlers struct {
	service service.ProdutoService
}

func (ch *ProdutoHandlers) GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["produto_id"]

	produto, err := ch.service.GetProduto(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, produto)
	}
}

func (ch *ProdutoHandlers) GetAllProduto(w http.ResponseWriter, r *http.Request) {
	produto, err := ch.service.GetAllProduto()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, produto)
	}
}

func (ch ProdutoHandlers) NewProduto(w http.ResponseWriter, r *http.Request) {
	var request dto.NewProdutoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		produto, appError := ch.service.NewProduto(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, produto)
		}
	}
}

func (ch ProdutoHandlers) UpdateProduto(w http.ResponseWriter, r *http.Request) {
	var request dto.NewProdutoResponse
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		produto, appError := ch.service.UpdateProduto(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, produto)
		}
	}
}

func (ch *ProdutoHandlers) DeleteProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["produto_id"]

	produto, err := ch.service.DeleteProduto(id)
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
