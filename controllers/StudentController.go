package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

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
	var updateRequest requests.Request[dto.Student]
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
		fieldsError, err := helpers.ValidateStruct(updateRequest.Data)

		if err != nil {
			response.StatusCode = http.StatusBadRequest
			response.Error = fieldsError
			response.Data = nil
			ReqAndRes.JSON(response.StatusCode, response)
			return
		} else {
			id, _ := updateRequest.TargetId.Int64()
			student, err := studentRepository.FindById(int(id))
			if err != nil {
				ReqAndRes.JSON(response.StatusCode, err)
				return
			}

			student.Nim = updateRequest.Data.Nim
			student.Name = updateRequest.Data.Name
			student.UpdatedAt = time.Now()

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

func (*StudentController) DeleteById(ReqAndRes *gin.Context) {
	studentRepository := ReqAndRes.MustGet("studentRepository").(*repositories.StudentRepository)
	response := &utils.ResponseMessage{Fields: helpers.GetFields(requests.DeleteRequest{})}
	deletedStudent := entities.Student{}
	deleteRequest := requests.DeleteRequest{}

	defer func() {
		if err := recover(); err != nil {
			response.StatusCode = http.StatusOK
			response.Error = err
			response.Data = nil
			ReqAndRes.JSON(response.StatusCode, response)
			return
		}
	}()

	err := ReqAndRes.ShouldBindJSON(&deleteRequest)
	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Error = err.Error()
		response.Data = nil
		ReqAndRes.JSON(response.StatusCode, response)
		return
	} else {
		fieldsError, err := helpers.ValidateStruct(deleteRequest)

		if err != nil {
			response.StatusCode = http.StatusBadRequest
			response.Error = fieldsError
			response.Data = nil
			ReqAndRes.JSON(response.StatusCode, response)
			return
		} else {
			id, _ := deleteRequest.TargetId.Int64()
			deletedStudent, err = studentRepository.DeleteById(int(id))
			if err != nil || deletedStudent.ID <= 0 {
				response.StatusCode = http.StatusBadRequest

				if deletedStudent.ID <= 0 {
					response.Error = errors.New("data not found").Error()
				} else {
					response.Error = err
				}

				response.Data = nil
				ReqAndRes.JSON(response.StatusCode, response)
				return
			} else {
				response.StatusCode = http.StatusOK
				response.Data = deletedStudent
				response.Error = nil
				ReqAndRes.JSON(response.StatusCode, response)
				return
			}
		}
	}
}
