package domain

import (
	"github.com/julianstephens/budget-tracker/cmd/postgresdb"
	"gorm.io/gorm"
)

const UniqueContstraintEmail = "idx_users_email"

type User struct {
	gorm.Model
	FullName  string
	ShortName string
	Email     string `gorm:"unique";not null`
	Password  string
}

type UserService interface {
	Validate(user *User) error
	GetAll(filters []postgresdb.DBFilter) ([]User, error)
	GetById(id uint) (User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type UserRepository interface {
	GetAll(filters []postgresdb.DBFilter) ([]User, error)
	GetById(id uint) (User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
