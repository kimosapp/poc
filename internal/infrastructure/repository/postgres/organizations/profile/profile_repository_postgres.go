package profile

import (
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewProfileRepositoryPostgres(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) GetAll() ([]organization.User, error) {
	var users []organization.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*organization.User, error) {
	var user organization.User
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *RepositoryPostgres) Create(user *organization.User) (
	*organization.User,
	error,
) {
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *RepositoryPostgres) Update(user *organization.User) (
	*organization.User,
	error,
) {
	if err := repo.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *RepositoryPostgres) Delete(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&organization.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}
