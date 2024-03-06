package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/dtos"
	"golang-restful-api/helper"
	"golang-restful-api/service"
	"net/http"
)

type CategoryController struct {
	Service service.CategoryService
}

func (controller *CategoryController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := dtos.CategoryRequestDtos{}

	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.POST(request.Context(), categoryCreateRequest)
	webResponse := dtos.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
