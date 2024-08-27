package router

import (
	"loantracker/deliverables/controller"
	"loantracker/infrastructure"

	"github.com/gin-gonic/gin"
)

// InitializeRouter sets up the router with all the endpoints
func InitializeRouter(usercontroller *controller.UserController) *gin.Engine {
	r := gin.Default()

	// Public endpoints
	r.POST("/register", usercontroller.Register)
	r.POST("/login", usercontroller.Login)
	r.POST("/reset-password", usercontroller.ResetPassword)

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(infrastructure.JWTMiddleware)
	protected.GET("/verify-email", usercontroller.VerifyEmail)
	protected.POST("/update-password", usercontroller.UpdatePassword)
	protected.GET("/users", usercontroller.GetUsersHandler)
	protected.GET("/users:id", usercontroller.GetUserHandler)
	protected.DELETE("/users/:id", usercontroller.DeleteUserHandler)
	return r
}
