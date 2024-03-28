package user

import (
	"errors"

	"dbo-be/entities"
	"dbo-be/helper"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SearchUsers(input UserSearchInput) (*helper.Pagination, error)
	GetUsers() (*helper.Pagination, error)
	RegisterUser(input RegisterUserInput) (entities.User, error)
	LoginUser(input LoginInput) (entities.User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUserByID(ID int) (entities.User, error)
	EditUser(ID string, input RegisterUserInput) (entities.User, error)
	DeleteUser(ID string) (entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SearchUsers(input UserSearchInput) (*helper.Pagination, error) {
	user, err := s.repository.SearchUser(input)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUsers() (*helper.Pagination, error) {
	users, err := s.repository.GetUser()

	if err != nil {
		return nil, err
	}

	return users, err
}

func (s *service) RegisterUser(input RegisterUserInput) (entities.User, error) {
	var user entities.User
	user.Name = input.Name
	user.NoHandphone = input.NoHandphone
	user.Email = input.Email
	user.Address = input.Address

	checkUserEmail, err := s.repository.FindUserByEmail(user.Email)
	if err != nil {
		return checkUserEmail, err
	}

	if checkUserEmail.ID != 0 {
		return checkUserEmail, errors.New("email ini sudah terdaftar")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (entities.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return false, err
	}

	// Jika email tidak ditemukan dalam database maka email available
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByID(ID int) (entities.User, error) {
	user, err := s.repository.FindUserByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan ID")
	}

	return user, nil
}

func (s *service) EditUser(ID string, input RegisterUserInput) (entities.User, error) {
	user, err := s.repository.UpdateUser(ID, input)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) DeleteUser(ID string) (entities.User, error) {
	user, err := s.repository.DestroyUser(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
