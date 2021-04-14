package repository

import (
    "github.com/angeldhakal/testcase-ms/models"
    "gorm.io/gorm"
)

type UserRepository interface {
	AddUser(models.Users) (models.Users, error)
	GetUser(int) (models.Users, error)
	GetUserByEmail(string) (models.Users, error)
	GetAllUsers() ([]models.Users, error)
	UpdateUser(models.Users) (models.Users, error)
	DeleteUser(models.Users) (models.Users, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: models.Connect(),
	}
}

func (db *userRepository) AddUser(user models.Users) (models.Users, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUser(id int) (user models.Users, err error) {
    return user, db.connection.First(&user, id).Error
}


func (db *userRepository) GetUserByEmail(email string) (user models.Users, err error) {
    return user, db.connection.First(&user, "email=?", email).Error
}

func (db *userRepository) GetAllUsers() (users []models.Users, err error) {
    return users, db.connection.Find(&users).Error
}

func (db *userRepository) UpdateUser(user models.Users) (models.Users, error) {
    if err := db.connection.First(&user, user.ID).Error; err != nil {
        return user, err
    }
    return user, db.connection.Model(&user).Updates(&user).Error
}

func (db *userRepository) DeleteUser(user models.Users) (models.Users, error) {
    if err := db.connection.First(&user, user.ID).Error; err != nil {
        return user, err
    }
    return user, db.connection.Model(&user).Error
}

