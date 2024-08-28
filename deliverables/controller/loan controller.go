package controller

import (
	"loantracker/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanController struct {
	loanUsecase domain.LoanUsecase
	logusecase  domain.LogUsecase
}

func NewLoanController(loanUsecase domain.LoanUsecase, logusecase domain.LogUsecase) *LoanController {
	return &LoanController{
		loanUsecase: loanUsecase,
		logusecase:  logusecase,
	}
}

func (lc *LoanController) Apply(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(domain.Claims)

	if claims.Role != "user" || claims.ID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	// Parse request body
	var loan domain.Loan

	if err := ctx.ShouldBindJSON(&loan); err != nil {
		return
	}
	if loan.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid loan amount"})
		return
	}
	userid, err := primitive.ObjectIDFromHex(claims.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	loan.UserID = userid
	loan.Status = "pending"

	// Call the loan use case to apply for a loan
	err = lc.loanUsecase.Apply(&loan)
	if err != nil {
		return
	}
	// Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Loan application submitted successfully",
		"loan_id": loan.ID.Hex(),
	})

	lc.logusecase.LogLoanApplication(loan.ID.Hex(), time.Now().Format("2006-01-02 15:04:05"))

}
func (lc *LoanController) View(ctx *gin.Context) {
	// Get the loan ID from the claims
	claims := ctx.MustGet("claims").(domain.Claims)
	if claims.Role != "user" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	loanid := ctx.Param("id")

	// Call the loan use case to retrieve the loan details
	loan, err := lc.loanUsecase.View(loanid)
	if err != nil {
		return
	}

	// Return the loan details as the response
	ctx.JSON(http.StatusOK, loan)

}
func (lc *LoanController) ViewAll(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(domain.Claims)

	if claims.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Call the loan use case to retrieve all loans
	loans, err := lc.loanUsecase.ViewAll()
	if err != nil {
		return
	}

	// Return the list of loans as the response
	ctx.JSON(http.StatusOK, loans)

}
func (lc *LoanController) ApproveReject(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(domain.Claims)

	if claims.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Get the loan ID from the request parameters
	loanID := ctx.Param("id")
	status := ctx.Param("status")

	if status != "approved" && status != "rejected" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid status"})
		return
	}

	// Call the loan use case to approve the loan
	err := lc.loanUsecase.ApproveReject(loanID, status)
	if err != nil {
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Loan updated successfully",
	})
	lc.logusecase.LogLoanApproval(loanID, status == "approved")

}
func (lc *LoanController) Delete(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(domain.Claims)
	if claims.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	// Get the loan ID from the request parameters
	loanID := ctx.Param("id")

	// Call the loan use case to delete the loan
	err := lc.loanUsecase.Delete(loanID)
	if err != nil {
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Loan deleted successfully",
	})
}
