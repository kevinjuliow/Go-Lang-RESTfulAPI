package middleware

import (
	"encoding/json"
	"golang-restful-api/helper"
	"golang-restful-api/models/dtos"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "TESTING" == request.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer ,request) //will continue to next Handler
	}else {
		//err
		writer.Header().Set("content-type" , "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := dtos.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED" ,
		}

		encoder := json.NewEncoder(writer)
		err := encoder.Encode(webResponse)
		helper.PanicIfError(err)
	}
}


