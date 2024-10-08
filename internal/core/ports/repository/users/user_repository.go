package users

import (
	"github.com/kimosapp/poc/internal/core/model/commons/types"
	"github.com/kimosapp/poc/internal/core/model/entity/users"
)

type UserRepository interface {
	GetAll() ([]users.User, error)
	GetPage(pageNumber int, pageSize int) (types.Page[users.User], error)
	GetByID(id string) (*users.User, error)
	GetByEmail(email string) (*users.User, error)
	Create(user *users.User) (*users.User, error)
	Update(user *users.User) (*users.User, error)
	Delete(id string) error
}
