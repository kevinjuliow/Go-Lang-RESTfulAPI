package ErrorHandler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"golang-restful-api/helper"
	"golang-restful-api/models/dtos"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter , request *http.Request , err interface{})  {
	if notFoundError(writer , request , err) {
		return
	}
	if validationError(writer , request , err) {
		return
	}
	internalServerError(writer , request , err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception , ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("content-type" , "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := dtos.WebResponse{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Data: exception.Error(),
		}
		encode := json.NewEncoder(writer)
		errEncode := encode.Encode(webResponse)
		helper.PanicIfError(errEncode)
		return true
	}else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception , ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("content-type" , "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := dtos.WebResponse{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Data: exception.Error,
		}
		encode := json.NewEncoder(writer)
		errEncode := encode.Encode(webResponse)
		helper.PanicIfError(errEncode)
		return true
	}else {
		return false
	}
}


func internalServerError(writer http.ResponseWriter , request *http.Request , err interface{}){
	writer.Header().Set("content-type" , "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := dtos.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "Internal Server Error" ,
		Data: err ,
	}

	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)
}