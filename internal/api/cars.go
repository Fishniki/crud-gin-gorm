package api

import (
	"context"
	"crudwebsocket/domain"
	"crudwebsocket/dto"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type carsApi struct {
	carsService domain.CarsService
	valiate     *validator.Validate
}

func NewCars(app *gin.Engine, carsService domain.CarsService, validate *validator.Validate) {
	ba := carsApi{
		carsService: carsService,
		valiate:     validate,
	}

	app.POST("/cars/create", ba.Create)
	app.GET("/cars/getall", ba.Index)
	app.GET("/cars/getbyid/:id", ba.Show)
	app.PUT("/cars/update/:id", ba.Update)
	app.DELETE("/cars/delet/:id", ba.Delete)

	app.Static("/media", "gallery")
}

func (ba carsApi) Create(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCarsRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.CreateResponsError("invalid request: "+err.Error()))
		return
	}

	file, err := ctx.FormFile("image")
	filename := uuid.NewString() + filepath.Ext(file.Filename)
	path := filepath.Join("gallery", filename)

	req.Image = filename

	err = ba.valiate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.CreateResponsError(err.Error()))
		return
	}

	err = ba.carsService.Create(c, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError(err.Error()))
		return
	}

	if err = ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError("failed to save image: "+ err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateResponsSucces(req))

}

func (ba carsApi) Index(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := ba.carsService.Index(c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError("Invalid Request"))
	}

	ctx.JSON(http.StatusOK, dto.CreateResponsSucces(res))
}

func (ba carsApi) Show(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Param("id")
	res, err := ba.carsService.Show(c, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateResponsSucces(res))
}

func (ba carsApi) Update(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCarsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.CreateResponsError("invalid request"))
		return
	}

	// ambil id dari URL
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.CreateResponsError("invalid uuid format"))
		return
	}

	// cek apakah data ada
	existingCar, err := ba.carsService.Show(c, id.String())
	if err != nil || existingCar.Id == uuid.Nil {
		ctx.JSON(http.StatusNotFound, dto.CreateResponsError("Data tidak ditemukan"))
		return
	}

	// set default Id dan Image lama
	req.Id = id
	req.Image = existingCar.Image

	// cek kalau ada file baru
	file, err := ctx.FormFile("image")
	if err == nil {
		// hapus file lama jika ada
		if existingCar.Image != "" {
			oldPath := filepath.Join("gallery", existingCar.Image)
			_ = os.Remove(oldPath)
		}

		// simpan file baru
		filename := uuid.NewString() + filepath.Ext(file.Filename)
		fmt.Println(filepath.Ext(file.Filename))
		newPath := filepath.Join("gallery", filename)
		if err := ctx.SaveUploadedFile(file, newPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError("failed to save image: "+err.Error()))
			return
		}
		req.Image = filename
	}

	// validasi struct (opsional)
	if err := ba.valiate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.CreateResponsError(err.Error()))
		return
	}

	// update data ke DB
	if err := ba.carsService.Update(c, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateResponsSucces("Data Berhasil di Update"))
}

func (ba carsApi) Delete(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Param("id")

	existingCar, err := ba.carsService.Show(c, id)

	if err != nil || existingCar.Id == uuid.Nil {
		ctx.JSON(http.StatusNotFound, dto.CreateResponsError("Id tidak di temukan " + err.Error()))
		return
	}

	if err := ba.carsService.Delete(c, id); err != nil{
		ctx.JSON(http.StatusInternalServerError, dto.CreateResponsError(err.Error()))
		return
	}
	
	
	if existingCar.Image != "" {
		oldPath := filepath.Join("gallery", existingCar.Image)
		_ = os.Remove(oldPath) // abaikan error kalau file tidak ada
	}


	ctx.JSON(http.StatusOK, dto.CreateResponsSucces("Data Berhasil di delete"))

}
