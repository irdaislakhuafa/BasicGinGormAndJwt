package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/database"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/helpers"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/repositories"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/routers"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/utils"
)

func main() {
	router := gin.Default()

	options := utils.OptionsFlags{}
	options.EnableOptionFlags()

	con := database.Connection{}
	con.SetOptionParams(&options)
	con.Setup()

	entities.Setup(&entities.Student{})

	router.Use(func(ctx *gin.Context) { // middleware
		studentRepository := &repositories.StudentRepository{}
		studentRepository.SetDB(database.GetDB())

		ctx.Set("studentRepository", studentRepository)
		ctx.Next()
	})

	// enable validation
	helpers.EnableValidator(true)

	v1 := routers.AppRouter{}
	v1.SetGroup(router.Group("/v1"))
	v1.Run()

	router.Run(":" + options.AppPort)
}
