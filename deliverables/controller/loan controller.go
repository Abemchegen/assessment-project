package controller

import "loantracker/usecase"

type LoanController struct {
	loanUsecase *usecase.LoanUsecase
}

func NewLoanController(loanUsecase *usecase.LoanUsecase) *LoanController {
	return &LoanController{
		loanUsecase: loanUsecase,
	}
}
