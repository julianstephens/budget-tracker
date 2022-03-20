package services

import (
	"errors"
	"github.com/julianstephens/budget-tracker/cmd/domain"
	"github.com/julianstephens/budget-tracker/cmd/postgresdb"
	"sync"
)

var once sync.Once

type userSVC struct {
	userRepo domain.UserRepository
}

var instance *userSVC

func CreateUserService(r domain.UserRepository) domain.UserService {
	once.Do(func() {
		instance = &userSVC{
			userRepo: r,
		}
	})

	return instance
}

func (*userSVC) Validate(user *domain.User) error {
	if user == nil {
		err := errors.New("user cannot be nil")
		return err
	}
	if user.FullName == "" {
		err := errors.New("user must have a full name")
		return err
	}
	if user.Email == "" {
		err := errors.New("user must have a registered email address")
		return err
	}

	return nil
}

func (u *userSVC) GetAll(filters []postgresdb.DBFilter) ([]domain.User, error) {
	return u.userRepo.GetAll(filters)
}

func (u *userSVC) GetById(id uint) (domain.User, error) {
	return u.userRepo.GetById(id)
}

func (u *userSVC) Create(user *domain.User) error {
	return u.userRepo.Create(user)
}

func (u *userSVC) Update(user *domain.User) error {
	return u.userRepo.Update(user)
}

func (u *userSVC) Delete(id uint) error {
	return u.userRepo.Delete(id)
}
