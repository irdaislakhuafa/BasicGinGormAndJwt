package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/database"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities"
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

	router.Run(":" + options.AppPort)
}
