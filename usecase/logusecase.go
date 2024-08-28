package usecase

import "loantracker/domain"

type LogUsecase struct {
	logRepo domain.LogRepository
}

func NewLogUsecase(logRepo domain.LogRepository) domain.LogUsecase {
	return &LogUsecase{
		logRepo: logRepo,
	}
}

func (lc *LogUsecase) ViewSystemLogs() (*domain.Logs, error) {
	return lc.logRepo.ViewSystemLogs()
}
func (lc *LogUsecase) LogLoginAttempt(email string, success bool, time string) error {
	return lc.logRepo.LogLoginAttempt(email, success, time)
}
func (lc *LogUsecase) LogPasswordResetRequest(email string, success bool, time string) error {
	return lc.logRepo.LogPasswordResetRequest(email, success, time)
}
func (lc *LogUsecase) LogPasswordResetSuccess(email string, success bool, time string) error {
	return lc.logRepo.LogPasswordResetSuccess(email, success, time)
}
func (lc *LogUsecase) LogLoanApplication(loanID string, time string) error {
	return lc.logRepo.LogLoanApplication(loanID, time)
}
func (lc *LogUsecase) LogLoanApproval(loanID string, status bool) error {
	return lc.logRepo.LogLoanApproval(loanID, status)
}
