package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities/dto"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/helpers"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/repositories"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/utils"
)

type StudentController struct {
}

func (s *StudentController) GetAll(ReqAndRes *gin.Context) {
	studentRepository := ReqAndRes.MustGet("studentRepository").(*repositories.StudentRepository)
	response := &utils.ResponseMessage{}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error :", r)
			response = &utils.ResponseMessage{
				StatusCode: http.StatusInternalServerError,
				Error:      r,
				Data:       nil,
			}
		}

		ReqAndRes.JSON(response.StatusCode, response)
	}()

	response = &utils.ResponseMessage{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       studentRepository.FindAll(),
	}
}

func (s *StudentController) Created(ReqAndRes *gin.Context) {
	studentRepository := ReqAndRes.MustGet("studentRepository").(*repositories.StudentRepository)

	response := &utils.ResponseMessage{}
	studentDto := dto.Student{}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error :", err)
			response = &utils.ResponseMessage{
				StatusCode: http.StatusInternalServerError,
				Error:      err,
				Data:       nil,
			}
			ReqAndRes.JSON(response.StatusCode, response)
		}
	}()

	err := ReqAndRes.ShouldBindJSON(&studentDto)
	studentDto.Nim = strings.Trim(studentDto.Nim, " ")
	studentDto.Name = strings.Trim(studentDto.Name, " ")

	if err != nil {
		response = &utils.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
		}
		ReqAndRes.JSON(response.StatusCode, response)
		return
	} else {

		fieldError, err := helpers.ValidateStruct(studentDto)

		if err != nil {
			response = &utils.ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Error:      fieldError,
				Data:       nil,
			}
			ReqAndRes.JSON(response.StatusCode, response)
			return
		} else {
			savedStudent, _ := studentRepository.Save(&entities.Student{Nim: studentDto.Nim, Name: studentDto.Name})
			response = &utils.ResponseMessage{
				StatusCode: http.StatusOK,
				Error:      nil,
				Data:       savedStudent,
			}
			ReqAndRes.JSON(response.StatusCode, response)
			return
		}
	}
}
