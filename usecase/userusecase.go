package usecase

import (
	"loantracker/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Register(user *domain.User) error {
	err := u.userRepo.Register(user)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) VerifyEmail(id string) error {
	err := u.userRepo.VerifyEmail(id)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) Login(user *domain.User) (string, error) {
	token, err := u.userRepo.Login(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (u *UserUsecase) ResetPassword(email string) error {
	err := u.userRepo.ResetPassword(email)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) UpdatePassword(email string, password string) error {
	err := u.userRepo.UpdatePassword(email, password)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) GetUsers() ([]domain.ResponceUser, error) {
	var user []domain.ResponceUser
	user, err := u.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *UserUsecase) GetUser(id string) (domain.ResponceUser, error) {
	var user domain.ResponceUser
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return domain.ResponceUser{}, err
	}
	return user, nil
}
func (u *UserUsecase) DeleteUser(id string) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
