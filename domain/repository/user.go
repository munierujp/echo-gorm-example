package repository

import (
	"echo-gorm-example/domain/model"
)

type UserRepository interface {
	Add(*model.User) error
	FindByID(uint) (*model.User, error)
	Last() (*model.User, error)
	Update(*model.User) error
	Delete(uint) error
}
