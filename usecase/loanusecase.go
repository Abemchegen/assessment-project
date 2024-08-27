package usecase

import "loantracker/repo"

type LoanUsecase struct {
	loanRepo *repo.LoanRepository
}

func NewLoanUsecase(loanRepo *repo.LoanRepository) *LoanUsecase {
	return &LoanUsecase{
		loanRepo: loanRepo,
	}
}
