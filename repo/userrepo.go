package repo

import (
	"context"
	"errors"
	"loantracker/domain"
	"loantracker/infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db         *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Client) domain.UserRepository {
	return &UserRepository{
		db:         db,
		database:   db.Database("Loan-tracker"),
		collection: db.Database("Loan-tracker").Collection("Loans"),
	}
}
func (r *UserRepository) Register(user *domain.User) error {
	user.Isverified = false
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	err = infrastructure.VerifyPasswordStrength(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	result, err := r.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	token, err := infrastructure.GenerateJWT(user.Name, user.ID.Hex(), user.Role, true)
	if err != nil {
		return err
	}
	err = infrastructure.SendVerificationEmail(user.Email, token)
	if err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) VerifyEmail(id string) error {

	userid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$set": bson.M{"isverified": true}})

	if err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) Login(user *domain.User) (string, error) {

	var User domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&User)
	if err != nil {
		return "", err
	}
	check := infrastructure.ComparePassword(user.Password, User.Password)
	if !check {
		return "", errors.New("invalid password")
	}
	if !User.Isverified {
		return "", errors.New("email not verified")
	}
	accesstoken, err := infrastructure.GenerateJWT(User.Name, User.ID.Hex(), User.Role, true)
	if err != nil {
		return "", err
	}
	refreshtoken, err := infrastructure.GenerateJWT(User.Name, User.ID.Hex(), User.Role, false)
	if err != nil {
		return "", err
	}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": User.ID}, bson.M{"$set": bson.M{"refreshtoken": refreshtoken}})
	if err != nil {
		return "", err
	}
	return accesstoken, nil
}
func (r *UserRepository) ResetPassword(email string) error {

	var User domain.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		return err
	}
	token, err := infrastructure.GenerateJWT(User.Name, User.ID.Hex(), User.Role, true)
	if err != nil {
		return err
	}
	err = infrastructure.SendResetEmail(email, token)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) UpdatePassword(email string, password string) error {
	err := infrastructure.VerifyPasswordStrength(password)
	if err != nil {
		return err
	}
	hashedPassword, err := infrastructure.HashPassword(password)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"email": email}, bson.M{"$set": bson.M{"password": hashedPassword}})
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}
func (r *UserRepository) GetUser(id string) (domain.User, error) {
	var user domain.User
	userid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": userid}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(id string) error {
	userid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": userid})
	if err != nil {
		return err
	}
	return nil
}
