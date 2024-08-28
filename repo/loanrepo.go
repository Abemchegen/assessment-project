package repo

import (
	"context"
	"loantracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (lr *LoanRepository) Apply(loan *domain.Loan) error {
	loan.ID = primitive.NewObjectID()
	_, err := lr.collection.InsertOne(context.TODO(), loan)
	if err != nil {
		return err
	}
	return nil
}
func (lr *LoanRepository) View(id string) (*domain.Loan, error) {
	var loan domain.Loan
	err := lr.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&loan)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}
func (lr *LoanRepository) ViewAll() ([]*domain.Loan, error) {
	var loans []*domain.Loan
	cursor, err := lr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var loan domain.Loan
		cursor.Decode(&loan)
		loans = append(loans, &loan)
	}
	return loans, nil
}
func (lr *LoanRepository) ApproveReject(id string, status string) error {

	_, err := lr.collection.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status}},
	)
	if err != nil {
		return err
	}
	return nil
}
func (lr *LoanRepository) Delete(id string) error {
	_, err := lr.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
