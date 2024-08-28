package domain

type Logs struct {
	LoginAttempts        []LoginAttempt         `json:"login_attempts,omitempty" bson:"login_attempts,omitempty"`
	PasswordResetRequest []PasswordResetRequest `json:"password_reset_requests,omitempty" bson:"password_reset_requests,omitempty"`
	LoanApplications     []LoanApplication      `json:"loan_applications,omitempty" bson:"loan_applications,omitempty"`
	LoanApprovals        []LoanApproval         `json:"loan_approvals,omitempty" bson:"loan_approvals,omitempty"`
}

type LoginAttempt struct {
	Email         string `json:"email,omitempty" bson:"email,omitempty"`
	Success       bool   `json:"success,omitempty" bson:"success,omitempty"`
	AttemptedTime string `json:"attemptedtime,omitempty" bson:"time,omitempty"`
}
type PasswordResetRequest struct {
	Email         string `json:"email,omitempty" bson:"email,omitempty"`
	Success       bool   `json:"success,omitempty" bson:"success,omitempty"`
	RequestedTime string `json:"requestedtime,omitempty" bson:"time,omitempty"`
	Successtime   string `json:"successtime,omitempty" bson:"time,omitempty"`
}
type LoanApplication struct {
	LoanID        string `json:"loan_id,omitempty" bson:"loan_id,omitempty"`
	SubmittedTime string `json:"submittedtime,omitempty" bson:"time,omitempty"`
}

type LoanApproval struct {
	LoanID string `json:"loan_id,omitempty" bson:"loan_id,omitempty"`
	Status bool   `json:"status,omitempty" bson:"status,omitempty"`
}

type LogRepository interface {
	LogLoginAttempt(email string, success bool, time string) error
	LogPasswordResetRequest(email string, success bool, time string) error
	LogPasswordResetSuccess(email string, success bool, time string) error
	LogLoanApplication(loanID string, time string) error
	LogLoanApproval(loanID string, status bool) error
	ViewSystemLogs() (*Logs, error)
}
type LogUsecase interface {
	LogLoginAttempt(email string, success bool, time string) error
	LogPasswordResetRequest(email string, success bool, time string) error
	LogPasswordResetSuccess(email string, success bool, time string) error
	LogLoanApplication(loanID string, time string) error
	LogLoanApproval(loanID string, status bool) error
	ViewSystemLogs() (*Logs, error)
}
