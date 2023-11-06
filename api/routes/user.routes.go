package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/api/controller"
)

// UserRoute -> Route for user module
type UserRoute struct {
	Handler    *gin.Engine
	Controller controller.UserController
}

// NewUserRoute -> initializes new instance of UserRoute
func NewUserRoute(
	controller controller.UserController,
	handler *gin.Engine,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// Setup -> setups user routes
func (u UserRoute) Setup() {
	user := u.Handler.Group("/api/v1/users")
	user.POST("/create", u.Controller.Create)
	user.PUT("/:id", u.Controller.Update)
	user.GET("/", u.Controller.ListAll)
	user.GET("/:id", u.Controller.GetById)
	user.DELETE("/:id", u.Controller.DeleteById)
}
