package services

import (
	"time"

	"github.com/ehardi19/bitment-backend/domain"

	"github.com/pkg/errors"
)

func (s *Services) GetAllUser() ([]domain.User, error) {
	return s.Repository.GetAllUser()
}

func (s *Services) CreateUser(user domain.User) (domain.User, int, error) {
	hash, err := s.Authentication.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, 400, err
	}

	user.ID = time.Now().UnixNano()
	user.Password = hash

	user, err = s.Repository.CreateUser(user)
	if err != nil {
		return domain.User{}, 400, err
	}

	accountBalance := domain.AccountBalance{
		UserID:  user.ID,
		Balance: 300000,
	}

	accountBalance, err = s.Repository.CreateAccountBalance(accountBalance)
	if err != nil {
		return domain.User{}, 400, err
	}

	accountCredit := domain.AccountCredit{
		UserID:  user.ID,
		Credits: 0,
	}

	accountCredit, err = s.Repository.CreateAccountCredit(accountCredit)
	if err != nil {
		return domain.User{}, 400, err
	}

	return user, 201, nil
}

func (s *Services) GetUserByID(id int64) (domain.User, error) {
	return s.Repository.GetUserByID(id)
}

func (s *Services) GetUserByEmail(email string) (domain.User, error) {
	return s.Repository.GetUserByEmail(email)
}

func (s *Services) UpdateUser(user domain.User) (domain.User, error) {
	_, err := s.Repository.GetUserByID(user.ID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	return user, s.Repository.UpdateUser(user)
}

func (s *Services) ChangePassword(user domain.User, req domain.ChangePasswordRequest) (domain.User, error) {
	userVerify, err := s.Repository.GetUserByID(user.ID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	if ok := s.Authentication.CheckPasswordHash(userVerify.Password, req.OldPassword); !ok {
		return domain.User{}, errors.New("old password does not match")
	}

	if req.OldPassword == req.NewPassword {
		return domain.User{}, errors.New("new password can not be same as old password")
	}

	hash, err := s.Authentication.HashPassword(req.NewPassword)
	if err != nil {
		return domain.User{}, errors.New("change password error")
	}

	user.Password = hash

	return user, s.Repository.UpdateUser(user)
}

func (s *Services) ChangePIN(user domain.User, req domain.ChangePINRequest) (domain.User, error) {
	userVerify, err := s.Repository.GetUserByID(user.ID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	if ok := s.Authentication.CheckPasswordHash(userVerify.PIN, req.OldPIN); !ok {
		return domain.User{}, errors.New("old password does not match")
	}

	if req.OldPIN == req.NewPIN {
		return domain.User{}, errors.New("new password can not be same as old password")
	}

	hash, err := s.Authentication.HashPassword(req.NewPIN)
	if err != nil {
		return domain.User{}, errors.New("change password error")
	}

	user.Password = hash

	return user, s.Repository.UpdateUser(user)
}
