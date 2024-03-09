package main

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"
)

func main() {
	db := app.NewDb()
	validator := validator.New()
	category_repository := repository.NewCategoryRepositoryImpl()
	category_service := service.NewCategoryServiceImpl(category_repository, db, validator)
	category_controller := controller.NewCategoryControllerImpl(category_service)

	router := httprouter.New()
	router.GET("/api/categories", category_controller.GetAll)
	router.GET("/api/categories/:category_id", category_controller.GetById)
	router.POST("/api/categories", category_controller.Create)
	router.PUT("/api/categories/:category_id", category_controller.Update)
	router.DELETE("/api/categories/:category_id", category_controller.Delete)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
