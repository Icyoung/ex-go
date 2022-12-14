package repo

import (
	"ex/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func User(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (repo *UserRepo) NewUser(name string, password string) (*model.User, error) {
	user := model.User{Name: name, Password: password}
	r := repo.DB.Model(&user).Create(&user)
	return &user, r.Error
}

func (repo *UserRepo) FindById(id uint) (*model.User, error) {
	user := model.User{}
	r := repo.DB.First(&user, id)
	return &user, r.Error
}

func (repo *UserRepo) FindByName(name string) (*model.User, error) {
	user := model.User{Name: name}
	r := repo.DB.Where(&user).First(&user)
	return &user, r.Error
}

func (repo *UserRepo) SaveToken(user *model.User, token string) *model.User {
	user.Token = token
	repo.DB.Save(&user)
	return user
}
