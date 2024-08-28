package controller

import (
	"loantracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	logUsecase domain.LogUsecase
}

func NewLogController(logUsecase domain.LogUsecase) *LogController {
	return &LogController{
		logUsecase: logUsecase,
	}
}

func (lc *LogController) ViewSystemLogs(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(domain.Claims)
	if claims.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Call the loan use case to retrieve system logs
	logs, err := lc.logUsecase.ViewSystemLogs()
	if err != nil {
		return
	}

	// Return the system logs as the response
	ctx.JSON(http.StatusOK, logs)

}
