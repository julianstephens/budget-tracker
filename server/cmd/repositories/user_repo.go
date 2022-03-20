package repositories

import (
	"fmt"
	"github.com/julianstephens/budget-tracker/cmd/domain"
	"github.com/julianstephens/budget-tracker/cmd/postgresdb"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type EmailDuplicateError struct {
	Email string
}

func CreateUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		DB: db,
	}
}

func (u *userRepo) GetAll(filters []postgresdb.DBFilter) ([]domain.User, error) {
	var users []domain.User
	if len(filters) == 0 {
		err := u.DB.Find(&users).Error
		return users, err
	}

	var query map[string]interface{}
	for _, f := range filters {
		query[f.FilterName] = f.FilterValue
	}
	err := u.DB.Where(query).Find(&users).Error
	return users, err
}

func (u *userRepo) GetById(id uint) (domain.User, error) {
	var user domain.User
	err := u.DB.First(&user, id).Error
	return user, err
}

func (u *userRepo) Create(user *domain.User) error {
	err := u.DB.Create(&user).Error
	if err != nil && postgresdb.IsUniqeConstraintError(err, domain.UniqueContstraintEmail) {
		return &EmailDuplicateError{Email: user.Email}
	}
	return err
}

func (u *userRepo) Update(user *domain.User) error {
	err := u.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Delete(id uint) error {
	err := u.DB.Delete(&domain.User{}, id).Error
	return err
}

func (e *EmailDuplicateError) Error() string {
	return fmt.Sprintf("Email '%s' already exists", e.Email)
}
