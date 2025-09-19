package main

import (
	"crudwebsocket/internal/api"
	"crudwebsocket/internal/config"
	"crudwebsocket/internal/connection"
	"crudwebsocket/internal/repository"
	"crudwebsocket/internal/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	cnf := config.Get()
	db, err := connection.GetDatabase(*cnf)
	validate := validator.New()

	if err != nil {
		log.Fatalf("Koneksi Gagal %s", err)
	}else{
		fmt.Println("Koneksi Berhasil")
	}

	app := gin.Default()

	carsRepsitory := repository.NewCars(db.GetDB())

	carsService := service.NewCars(carsRepsitory)

	api.NewCars(app, carsService, validate)

	app.Run()
}