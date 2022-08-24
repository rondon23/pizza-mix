package app

import (
	"encoding/json"
	"net/http"
	"pizza-backend/dto"
	"pizza-backend/service"
)

type ClienteHandlers struct {
	service service.ClienteService
}

func (ch *ClienteHandlers) GetAllClientes(w http.ResponseWriter, r *http.Request) {
	cliente, err := ch.service.GetAllCliente()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, cliente)
	}
}

func (ch ClienteHandlers) NewCliente(w http.ResponseWriter, r *http.Request) {
	var request dto.NewClienteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		cliente, appError := ch.service.NewCliente(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, cliente)
		}
	}
}
