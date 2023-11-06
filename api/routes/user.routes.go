package routes

import (
	"go-crud/api/controller"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controller.UserController
}

func NewUserRouteController(
	userController controller.UserController,
) UserRouteController {
	return UserRouteController{userController}
}

func (u *UserRouteController) Setup(r *gin.RouterGroup) {
	r.POST("/create", u.userController.Create)
	r.GET("/:id", u.userController.GetById)
	r.GET("/", u.userController.ListAll)
	r.PUT("/:id", u.userController.Update)
	r.DELETE("/:id", u.userController.DeleteById)
}
