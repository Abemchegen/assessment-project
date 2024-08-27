package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	Isverified   bool               `json:"isverified,omitempty" bson:"isverified,omitempty"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	Loans        []Loan             `json:"loans,omitempty" bson:"loans,omitempty"`
	RefreshToken string             `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
}

type ResponceUser struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Isverified bool               `json:"isverified,omitempty" bson:"isverified,omitempty"`
	Role       string             `json:"role,omitempty" bson:"role,omitempty"`
	Loans      []Loan             `json:"loans,omitempty" bson:"loans,omitempty"`
}

type UserRepository interface {
	Register(user *User) error
	VerifyEmail(id string) error
	Login(user *User) (string, error)
	ResetPassword(email string) error
	UpdatePassword(token string, password string) error
	GetUsers() ([]ResponceUser, error)
	GetUser(id string) (ResponceUser, error)
	DeleteUser(id string) error
}
type UserUsecase interface {
	Register(user *User) error
	VerifyEmail(id string) error
	Login(user *User) (string, error)
	ResetPassword(email string) error
	UpdatePassword(token string, password string) error
	GetUsers() ([]ResponceUser, error)
	GetUser(id string) (ResponceUser, error)
	DeleteUser(id string) error
}
