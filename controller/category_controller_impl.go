package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/helper"
	dtos2 "golang-restful-api/models/dtos"
	"golang-restful-api/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryControllerImpl(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{Service: service}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//Read json from request body
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := dtos2.CategoryCreateDtos{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	//Call the POST from service and convert to webResponse
	categoryResponse := controller.Service.POST(request.Context(), categoryCreateRequest)
	fmt.Println("in controller")
	fmt.Println(categoryResponse.Name)
	fmt.Println(categoryResponse.Id)
	webResponse := dtos2.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   categoryResponse,
	}

	//Add Header
	writer.Header().Add("content-type", "application/json")

	//Encode the WebResponse into JSON directly in the writer
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryRequest := dtos2.CategoryUpdateDtos{}
	err := decoder.Decode(&categoryRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("category_id")
	id, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	categoryRequest.Id = uint(id)

	writer.Header().Add("content-type", "application/json")
	categoryResponse := controller.Service.PUT(request.Context(), categoryRequest)
	webResponse := dtos2.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("category_id")

	parsedId, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	controller.Service.DELETE(request.Context(), uint(parsedId))
	webResponse := dtos2.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("category_id")
	parsedId, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.GETById(request.Context(), uint(parsedId))

	webResponse := dtos2.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryLists := controller.Service.GETAll(request.Context())

	webResponse := dtos2.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryLists,
	}

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
