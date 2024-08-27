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
	r.POST("/resetpassword", usercontroller.ResetPassword)

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(infrastructure.JWTMiddleware)
	protected.GET("/verifyemail", usercontroller.VerifyEmail)
	protected.POST("/updatepassword", usercontroller.UpdatePassword)
	protected.GET("/users", usercontroller.GetUsersHandler)
	protected.GET("/user", usercontroller.GetUserHandler)
	protected.DELETE("/users/:id", usercontroller.DeleteUserHandler)
	return r
}
