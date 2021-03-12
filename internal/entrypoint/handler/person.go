package handler

import (
	"jeeo/api-boilerplate/internal/application/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PersonHandler struct {
	service *service.PersonService
}

// func (ph *PersonHandler) Create(payload dto.PersonDTO) dto.PersonDTO {
// 	person := ph.repository.Create(ph.mapper.FromDTO(payload))

// 	return ph.mapper.ToDTO(person)
// }

func (ph *PersonHandler) GetById(ctx echo.Context) error {
	personId, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println("error on parse HTTP param: ", err.Error())
		return err
	}
	response := ph.service.GetById(int32(personId))

	ctx.JSON(http.StatusOK, response)
	return nil
}

// func (ph *PersonHandler) GetAll() []dto.PersonDTO {
// 	people := ph.repository.GetAll()

// 	return ph.mapper.ToManyDTO(people)
// }

// func (ph *PersonHandler) Update(payload dto.PersonDTO) dto.PersonDTO {
// 	person := ph.repository.Update(ph.mapper.FromDTO(payload))

// 	return ph.mapper.ToDTO(person)
// }

// func (ph *PersonHandler) Delete(id int32) bool {
// 	success := ph.repository.Delete(id)

// 	return success
// }

func ProvidePersonHandler(s *service.PersonService) *PersonHandler {
	return &PersonHandler{
		service: s,
	}
}
