package handler

import (
	"encoding/json"
	"net/http"
	"stakeholders/model"
	"stakeholders/service"
)

type AppRatingHandler struct {
	AppRatingService *service.AppRatingService
}

func (handler *AppRatingHandler) Create(writer http.ResponseWriter, req *http.Request) {

	var appRating model.AppRating
	err := json.NewDecoder(req.Body).Decode(&appRating)

	if err != nil {
		println("Error while parsing ", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdApp, err := handler.AppRatingService.Create(&appRating)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(createdApp)
}

func (handler *AppRatingHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	ratings, err := handler.AppRatingService.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(ratings)
}
