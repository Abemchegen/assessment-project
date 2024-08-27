package controller

import (
	"loantracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.UserUsecase
	logusecase  domain.LogUsecase
}

func NewUserController(userUsecase domain.UserUsecase, logusecase domain.LogUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		logusecase:  logusecase,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user domain.User

	// Bind the request body to the user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name, email and password are required"})
		return
	}

	// Call the user usecase to register the user
	err := c.userUsecase.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully, verification email sent"})
}
func (c *UserController) VerifyEmail(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(*domain.Claims)
	if claims.Name == "" || claims.ID == "" || claims.Role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Verification code is required"})
		return
	}

	// Verify the email using the verification code
	err := c.userUsecase.VerifyEmail(claims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})

}
func (c *UserController) Login(ctx *gin.Context) {
	var user domain.User

	// Bind the request body to the loginRequest struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Call the user usecase to authenticate the user
	token, err := c.userUsecase.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
func (c *UserController) ResetPassword(ctx *gin.Context) {
	var resetRequest struct {
		Email string `json:"email"`
	}

	// Bind the request body to the resetRequest struct
	if err := ctx.ShouldBindJSON(&resetRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if resetRequest.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Call the user usecase to initiate the password reset process
	err := c.userUsecase.ResetPassword(resetRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset initiated, check your email for instructions"})
}
func (c *UserController) UpdatePassword(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(*domain.Claims)

	if claims.Name == "" || claims.ID == "" || claims.Role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}

	var updateRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the request body to the updateRequest struct
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updateRequest.Email == "" || updateRequest.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and new password are required"})
		return
	}

	// Call the user usecase to update the user's password
	err := c.userUsecase.UpdatePassword(updateRequest.Email, updateRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
func (c *UserController) GetUsersHandler(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.Claims)
	if claims.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	users, err := c.userUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
func (c *UserController) GetUserHandler(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.Claims)
	if (claims.Role != "user" && claims.Role != "admin") || claims.ID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Call the user usecase to get the user
	user, err := c.userUsecase.GetUser(claims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
func (c *UserController) DeleteUserHandler(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*domain.Claims)
	if claims.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user ID from the request query parameter
	userID := ctx.Query("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Call the user usecase to delete the user
	err := c.userUsecase.DeleteUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
