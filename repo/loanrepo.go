package repo

import (
	"loantracker/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoanRepository struct {
	db         *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewLoanRepository(db *mongo.Client) domain.LoanRepository {
	return &LoanRepository{
		db:         db,
		database:   db.Database("Loan-tracker"),
		collection: db.Database("Loan-tracker").Collection("Loans"),
	}
}
