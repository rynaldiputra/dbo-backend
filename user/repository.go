package user

import (
	"dbo-be/entities"
	"dbo-be/helper"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	SearchUser(input UserSearchInput) (*helper.Pagination, error)
	GetUser() (*helper.Pagination, error)
	CreateUser(entities.User) (entities.User, error)
	FindUserByEmail(email string) (entities.User, error)
	FindUserByID(ID int) (entities.User, error)
	UpdateUser(ID string, input RegisterUserInput) (entities.User, error)
	DestroyUser(ID string) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SearchUser(input UserSearchInput) (*helper.Pagination, error) {
	var users []*entities.User
	var pagination helper.Pagination
	pagination.Sort = "id ASC"
	condition := fmt.Sprintf("%s LIKE ?", input.Column)
	value := "%" + input.Value + "%"

	r.db.Scopes(helper.Paginate(users, &pagination, r.db)).Where(condition, value).Find(&users)

	pagination.Rows = users

	return &pagination, nil
}

func (r *repository) GetUser() (*helper.Pagination, error) {
	var users []*entities.Order
	var pagination helper.Pagination
	pagination.Sort = "id ASC"

	r.db.Scopes(helper.Paginate(users, &pagination, r.db)).Preload("Orders").Find(&users)

	pagination.Rows = users

	return &pagination, nil
}

func (r *repository) CreateUser(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByEmail(email string) (entities.User, error) {
	var user entities.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByID(ID int) (entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(ID string, input RegisterUserInput) (entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return user, errors.New("Data User tidak ditemukan")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	user.Name = input.Name
	user.NoHandphone = input.NoHandphone
	user.Email = input.Email
	user.Password = string(passwordHash)
	user.Address = input.Address
	user.UpdatedAt = time.Now()

	err = r.db.Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DestroyUser(ID string) (entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return user, errors.New("Data user tidak ditemukan")
	}

	err = r.db.Delete(&user, ID).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
