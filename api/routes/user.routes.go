package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/api/controller"
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
	router := r.Group("/api/v1/users")
	router.POST("/create", u.userController.Create)
	router.GET("/:id", u.userController.GetById)
	router.GET("/", u.userController.ListAll)
	router.PUT("/:id", u.userController.Update)
	router.DELETE("/:id", u.userController.DeleteById)
}
