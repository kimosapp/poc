package users

import (
	"github.com/kimosapp/poc/internal/core/model/commons/types"
	userEntity "github.com/kimosapp/poc/internal/core/model/entity/users"
)

type Repository interface {
	GetAll() ([]userEntity.User, error)
	GetPage(pageNumber int, pageSize int) (types.Page[userEntity.User], error)
	GetByID(id string) (*userEntity.User, error)
	GetByEmail(email string) (*userEntity.User, error)
	Create(user *userEntity.User) (*userEntity.User, error)
	Update(user *userEntity.User) (*userEntity.User, error)
	Delete(id string) error
	GetAllByEmail(emails []string) ([]userEntity.User, error)
}
