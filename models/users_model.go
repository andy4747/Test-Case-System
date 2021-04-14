package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	Superuser bool   `json:"superuser"`
}

func (Users) TableName() string {
	return "users"
}

type UserRepository interface {
	AddUser(Users) (Users, error)
	GetUser(int) (Users, error)
	GetUserByEmail(string) (Users, error)
	GetAllUsers() ([]Users, error)
	UpdateUser(Users) (Users, error)
	DeleteUser(Users) (Users, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: Connect(),
	}
}

func (db *userRepository) AddUser(user Users) (Users, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUser(id int) (user Users, err error) {
    return user, db.connection.First(&user, id).Error
}


func (db *userRepository) GetUserByEmail(email string) (user Users, err error) {
    return user, db.connection.First(&user, "email=?", email).Error
}

func (db *userRepository) GetAllUsers() (users []Users, err error) {
    return users, db.connection.Find(&users).Error
}

func (db *userRepository) UpdateUser(user Users) (Users, error) {
    if err := db.connection.First(&user, user.ID).Error; err != nil {
        return user, err
    }
    return user, db.connection.Model(&user).Updates(&user).Error
}

func (db *userRepository) DeleteUser(user Users) (Users, error) {
    if err := db.connection.First(&user, user.ID).Error; err != nil {
        return user, err
    }
    return user, db.connection.Model(&user).Error
}

