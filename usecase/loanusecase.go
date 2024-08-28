package usecase

import (
	"loantracker/domain"
)

type LoanUsecase struct {
	loanRepo domain.LoanRepository
}

func NewLoanUsecase(loanRepo domain.LoanRepository) domain.LoanUsecase {
	return &LoanUsecase{
		loanRepo: loanRepo,
	}
}

func (lc *LoanUsecase) Apply(loan *domain.Loan) error {
	return lc.loanRepo.Apply(loan)
}
func (lc *LoanUsecase) View(id string) (*domain.Loan, error) {
	return lc.loanRepo.View(id)
}
func (lc *LoanUsecase) ViewAll() ([]*domain.Loan, error) {
	return lc.loanRepo.ViewAll()
}
func (lc *LoanUsecase) ApproveReject(id string, status string) error {
	return lc.loanRepo.ApproveReject(id, status)
}
func (lc *LoanUsecase) Delete(id string) error {
	return lc.loanRepo.Delete(id)
}
