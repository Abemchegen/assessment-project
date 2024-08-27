package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id, omitempty"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id, omitempty"`
	Amount float64            `json:"amount,omitempty" bson:"amount, omitempty"`
	Status string             `json:"status,omitempty" bson:"status, omitempty"`
}

type LoanRepository interface {
}
type LoanUsecase interface {
}
