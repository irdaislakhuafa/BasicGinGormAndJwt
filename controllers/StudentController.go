package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities/dto"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/helpers"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/repositories"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/utils"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/utils/requests"
)

type StudentController struct {
}

func (*StudentController) GetAll(ReqAndRes *gin.Context) {
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

func (*StudentController) Created(ReqAndRes *gin.Context) {
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
			savedStudent, err := studentRepository.Save(&entities.Student{Nim: studentDto.Nim, Name: studentDto.Name})

			if err != nil {
				response = &utils.ResponseMessage{
					StatusCode: http.StatusBadRequest,
					Error:      err.Error(),
					Data:       nil,
				}
				ReqAndRes.JSON(response.StatusCode, response)
				return
			}
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

func (*StudentController) UpdateById(ReqAndRes *gin.Context) {
	studentRepository := ReqAndRes.MustGet("studentRepository").(*repositories.StudentRepository)
	var updateRequest requests.UpdateRequest[dto.Student]
	response := &utils.ResponseMessage{}
	response.Fields = helpers.GetFields(updateRequest)

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error :", r)
			response.StatusCode = http.StatusInternalServerError
			response.Error = r
			response.Data = nil
			ReqAndRes.JSON(response.StatusCode, response)
		}
	}()

	err := ReqAndRes.ShouldBindJSON(&updateRequest)

	if err != nil {
		response = &utils.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
		}
		ReqAndRes.JSON(response.StatusCode, response)
		return
	} else {
		fieldsError, err := helpers.ValidateStruct(updateRequest.NewData)

		if err != nil {
			response.StatusCode = http.StatusBadRequest
			response.Error = fieldsError
			response.Data = nil
			ReqAndRes.JSON(response.StatusCode, response)
			return
		} else {

			student := entities.Student{
				ID:   uint64(updateRequest.TargetId),
				Nim:  updateRequest.NewData.Nim,
				Name: updateRequest.NewData.Name,
			}

			student, err = studentRepository.Save(&student)

			if err != nil {
				response.StatusCode = http.StatusBadRequest
				response.Error = err.Error()
				response.Data = nil
				ReqAndRes.JSON(response.StatusCode, response)
				return
			} else {
				response.StatusCode = http.StatusOK
				response.Error = nil
				response.Data = student
				ReqAndRes.JSON(response.StatusCode, response)
				return
			}
		}
	}
}
