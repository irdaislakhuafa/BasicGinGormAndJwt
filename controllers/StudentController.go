package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities"
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

func (s *StudentController) Created(ReqAndReq *gin.Context) {
	response := &utils.ResponseMessage{}
	student := entities.Student{}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error :", err)
			response = &utils.ResponseMessage{
				StatusCode: http.StatusInternalServerError,
				Error:      err,
				Data:       nil,
			}
		}
		ReqAndReq.JSON(response.StatusCode, response)
	}()

	err := ReqAndReq.BindJSON(&student)
	if err != nil {
		response = &utils.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
		}
	} else {
		response = &utils.ResponseMessage{
			StatusCode: http.StatusOK,
			Error:      nil,
			Data:       student,
		}
	}

}
