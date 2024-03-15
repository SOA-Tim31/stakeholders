package handler

import (
	"encoding/json"
	"net/http"
	"stakeholders/model"
	"stakeholders/service"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	PersonService *service.PersonService
}

func (handler *PersonHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	person, err := handler.PersonService.FindPerson(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(person)
}

func (handler *PersonHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var person model.Person
	err := json.NewDecoder(req.Body).Decode(&person)

	if err != nil {
		println("Error while parsing ", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedPerson, err2 := handler.PersonService.Update(&person)
	if err2 != nil {
		println("Error while updateing a new user")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedPerson)
}

func (handler *PersonHandler) Create(writer http.ResponseWriter, req *http.Request) {

}
