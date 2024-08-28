package router

import (
	"loantracker/deliverables/controller"
	"loantracker/infrastructure"

	"github.com/gin-gonic/gin"
)

// InitializeRouter sets up the router with all the endpoints
func InitializeRouter(usercontroller *controller.UserController, loancontroller *controller.LoanController, logcontroller *controller.LogController) *gin.Engine {
	r := gin.Default()

	// Public endpoints
	r.POST("/register", usercontroller.Register)
	r.POST("/login", usercontroller.Login)
	r.POST("/resetpassword", usercontroller.ResetPassword)

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(infrastructure.JWTMiddleware)

	// user endpoints
	protected.GET("/verifyemail", usercontroller.VerifyEmail)
	protected.POST("/updatepassword", usercontroller.UpdatePassword)
	protected.GET("/users", usercontroller.GetUsersHandler)
	protected.GET("/user", usercontroller.GetUserHandler)
	protected.DELETE("/users/:id", usercontroller.DeleteUserHandler)

	// Loan endpoints
	protected.POST("/apply", loancontroller.Apply)
	protected.GET("/view/:id", loancontroller.View)
	protected.GET("/view", loancontroller.ViewAll)
	protected.DELETE("/delete/:id", loancontroller.Delete)
	protected.PUT("/update/:id", loancontroller.ApproveReject)

	// Log endpoints
	protected.GET("/logs", logcontroller.ViewSystemLogs)

	return r
}
