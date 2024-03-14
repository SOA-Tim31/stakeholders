package handler

import (
	"net/http"
	"stakeholders/service"
)

type PersonHandler struct {
	PersonService *service.PersonService
}

func (handler *PersonHandler) Get(writer http.ResponseWriter, req *http.Request) {

}

func (handler *PersonHandler) Create(writer http.ResponseWriter, req *http.Request) {

}
