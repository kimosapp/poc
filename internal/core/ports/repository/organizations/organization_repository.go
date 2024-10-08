package organization

import (
	"github.com/kimosapp/poc/internal/core/model/commons/types"
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]organization.Organization, error)
	GetPage(pageNumber int, pageSize int) (types.Page[organization.Organization], error)
	GetByID(id string) (*organization.Organization, error)
	GetByIDAndUserId(orgId string, userId string) (*organization.Organization, error)
	GetAllByUserId(userId string) ([]organization.Organization, error)
	Create(organization *organization.Organization, tx *gorm.DB) (*organization.Organization, error)
	Update(organization *organization.Organization) (*organization.Organization, error)
	Delete(id string) error
	BeginTransaction() *gorm.DB
}
