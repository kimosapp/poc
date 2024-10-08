package team

import (
	"github.com/kimosapp/poc/internal/core/model/commons/types"
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllByOrgId(orgId string) ([]organization.Team, error)
	GetPageByOrgId(orgId string, page int, size int) (types.Page[organization.Team], error)
	GetByID(id string) (*organization.Team, error)
	GetByNameOrSlugAndOrgId(name string, slug string, id string) ([]organization.Team, error)
	Create(team *organization.Team, tx *gorm.DB) (*organization.Team, error)
	Update(team *organization.Team, tx *gorm.DB) (*organization.Team, error)
	Delete(id string, tx *gorm.DB) error
	BeginTransaction() *gorm.DB
}
