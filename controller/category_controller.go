package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/dtos"
	"golang-restful-api/helper"
	"golang-restful-api/service"
	"net/http"
	"strconv"
)

type CategoryController struct {
	Service service.CategoryService
}

func (controller *CategoryController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//Read json from request body
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := dtos.CategoryRequestDtos{}

	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.POST(request.Context(), categoryCreateRequest)
	webResponse := dtos.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   categoryResponse,
	}
	writer.Header().Add("content-type", "application/json")
	//Encode the WebResponse into JSON directly in the writer
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryRequest := dtos.CategoryRequestDtos{}
	err := decoder.Decode(&categoryRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("category_id")
	id, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	categoryRequest.Id = uint(id)

	writer.Header().Add("content-type", "application/json")
	categoryResponse := controller.Service.PUT(request.Context(), categoryRequest)
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

func (controller *CategoryController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("category_id")

	parsedId, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	controller.Service.DELETE(request.Context(), uint(parsedId))
	webResponse := dtos.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("category_id")
	parsedId, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.GETById(request.Context(), uint(parsedId))

	webResponse := dtos.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryLists := controller.Service.GETAll(request.Context())

	webResponse := dtos.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryLists,
	}

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
