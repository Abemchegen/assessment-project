package repo

import (
	"context"
	"loantracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository struct {
	db                             *mongo.Client
	database                       *mongo.Database
	loginattemptcollection         *mongo.Collection
	passwordresetrequestcollection *mongo.Collection
	loanapplicationcollection      *mongo.Collection
	loanapprovalcollection         *mongo.Collection
}

func NewLogRepository(db *mongo.Client) domain.LogRepository {
	return &LogRepository{
		db:                             db,
		database:                       db.Database("Loan-tracker"),
		loginattemptcollection:         db.Database("Loan-tracker").Collection("loginattemptlogs"),
		passwordresetrequestcollection: db.Database("Loan-tracker").Collection("passwordresetrequestlogs"),
		loanapplicationcollection:      db.Database("Loan-tracker").Collection("loanapplicationslogs"),
		loanapprovalcollection:         db.Database("Loan-tracker").Collection("loanapprovallogs"),
	}
}

func (lr *LogRepository) LogLoginAttempt(email string, success bool, time string) error {
	_, err := lr.loginattemptcollection.InsertOne(context.TODO(), bson.M{"email": email, "success": success, "time": time})
	if err != nil {
		return err
	}
	return nil
}
func (lr *LogRepository) LogPasswordResetRequest(email string, success bool, time string) error {
	_, err := lr.passwordresetrequestcollection.InsertOne(context.TODO(), bson.M{"email": email, "success": success, "time": time})
	if err != nil {
		return err
	}
	return nil
}
func (lr *LogRepository) LogPasswordResetSuccess(email string, success bool, time string) error {
	_, err := lr.passwordresetrequestcollection.UpdateOne(context.TODO(), bson.M{"email": email, "success": success, "time": time}, bson.M{"$set": bson.M{"email": email, "success": success, "time": time}})
	if err != nil {
		return err
	}

	return nil

}
func (lr *LogRepository) LogLoanApplication(loanID string, time string) error {
	_, err := lr.loanapplicationcollection.InsertOne(context.TODO(), bson.M{"loan_id": loanID, "time": time})
	if err != nil {
		return err
	}
	return nil
}

func (lr *LogRepository) LogLoanApproval(loanID string, status bool) error {
	_, err := lr.loanapplicationcollection.UpdateOne(context.TODO(), bson.M{"loan_id": loanID}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return err
	}
	return nil

}

func (lr *LogRepository) ViewSystemLogs() (*domain.Logs, error) {
	logs := domain.Logs{} // Initialize logs with an empty value

	var LoginAttempts []domain.LoginAttempt
	var PasswordResetRequest []domain.PasswordResetRequest
	var LoanApplications []domain.LoanApplication
	var loanApprovals []domain.LoanApproval

	cursor, err := lr.loginattemptcollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return &logs, err // Return logs instead of nil
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var log domain.LoginAttempt
		err := cursor.Decode(&log)
		if err != nil {
			return &logs, err // Return logs instead of nil
		}
		LoginAttempts = append(LoginAttempts, log)
	}

	cursor, err = lr.passwordresetrequestcollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return &logs, err // Return logs instead of nil
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var log domain.PasswordResetRequest
		err := cursor.Decode(&log)
		if err != nil {
			return &logs, err // Return logs instead of nil
		}
		PasswordResetRequest = append(PasswordResetRequest, log)
	}

	cursor, err = lr.loanapplicationcollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return &logs, err // Return logs instead of nil
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var log domain.LoanApplication
		err := cursor.Decode(&log)
		if err != nil {
			return &logs, err // Return logs instead of nil
		}
		LoanApplications = append(LoanApplications, log)
	}

	cursor, err = lr.loanapprovalcollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return &logs, err // Return logs instead of nil
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var log domain.LoanApproval
		err := cursor.Decode(&log)
		if err != nil {
			return &logs, err // Return logs instead of nil
		}
		loanApprovals = append(loanApprovals, log)
	}

	logs.LoginAttempts = LoginAttempts
	logs.PasswordResetRequest = PasswordResetRequest
	logs.LoanApplications = LoanApplications
	logs.LoanApprovals = loanApprovals

	return &logs, nil
}
