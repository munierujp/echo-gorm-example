package database

import (
	"echo-gorm-example/domain/model"
	"echo-gorm-example/domain/repository"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Add(user model.User) error {
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Last() (*model.User, error) {
	var user model.User
	if err := repo.db.Last(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *model.User) error {
	if err := repo.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(id uint) error {
	user := model.User{}
	user.ID = id
	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
