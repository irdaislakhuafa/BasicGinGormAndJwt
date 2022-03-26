package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/controllers"
)

type AppRouter struct {
	group *gin.RouterGroup
}

func (app *AppRouter) Run() {
	studentController := controllers.StudentController{}

	app.group.GET("/students", studentController.GetAll)
	app.group.POST("/students", studentController.Created)
	app.group.PUT("/students", studentController.UpdateById)
	app.group.DELETE("/students", studentController.DeleteById)
}
func (app *AppRouter) SetGroup(group *gin.RouterGroup) {
	app.group = group
}
