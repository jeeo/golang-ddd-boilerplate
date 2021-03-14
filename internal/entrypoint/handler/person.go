package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Jeeo/golang-ddd-boilerplate/internal/application"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"

	"github.com/labstack/echo/v4"
)

type PersonHandlerImpl struct {
	application application.PersonApplication
}

func (ph *PersonHandlerImpl) Create(ctx echo.Context) error {
	personDTO := new(dto.PersonDTO)
	if err := ctx.Bind(personDTO); err != nil {
		log.Println("error on parse body: ", err.Error())
		ctx.Response().Status = 500
		return err
	}
	response := ph.application.CreatePerson(*personDTO)
	ctx.JSON(http.StatusOK, response)
	return nil
}

func (ph *PersonHandlerImpl) GetById(ctx echo.Context) error {
	personId, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println("error on parse HTTP param: ", err.Error())
		ctx.Response().Status = 500
		return err
	}
	response := ph.application.FindOne(int32(personId))

	ctx.JSON(http.StatusOK, response)
	return nil
}

func (ph *PersonHandlerImpl) GetAll(ctx echo.Context) error {
	response := ph.application.FindAll()

	ctx.JSON(http.StatusOK, response)
	return nil
}

func (ph *PersonHandlerImpl) Update(ctx echo.Context) error {
	personId, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println("error on parse HTTP param: ", err.Error())
		ctx.Response().Status = 500
		return err
	}
	personDTO := new(dto.PersonDTO)
	if err = ctx.Bind(personDTO); err != nil {
		log.Println("error on parse body: ", err.Error())
		ctx.Response().Status = 500
		return err
	}
	response := ph.application.UpdatePerson(int32(personId), *personDTO)
	ctx.JSON(http.StatusOK, response)
	return nil
}

func (ph *PersonHandlerImpl) Delete(ctx echo.Context) error {
	personId, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println("error on parse HTTP param: ", err.Error())
		return err
	}

	response := ph.application.DeletePerson(int32(personId))
	ctx.JSON(http.StatusOK, response)
	return nil
}

func ProvidePersonHandler(a application.PersonApplication) *PersonHandlerImpl {
	return &PersonHandlerImpl{
		application: a,
	}
}
