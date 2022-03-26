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
}
func (app *AppRouter) SetGroup(group *gin.RouterGroup) {
	app.group = group
}
