package postgres

import (
	types "github.com/kimosapp/poc/internal/core/model/commons/types"
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	organization2 "github.com/kimosapp/poc/internal/core/ports/repository/organizations"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) organization2.Repository {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) GetAll() ([]organization.Organization, error) {
	var organizations []organization.Organization
	if err := repo.db.Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

func (repo *RepositoryPostgres) GetPage(
	pageNumber int,
	pageSize int,
) (types.Page[organization.Organization], error) {
	var totalRows int64
	repo.db.Model(&organization.Organization{}).Count(&totalRows)
	var organizations []organization.Organization
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&organizations).Error; err != nil {
		return types.EmptyPage[organization.Organization](), err
	}
	pageBuilder := new(types.PageBuilder[organization.Organization])
	return pageBuilder.SetItems(organizations).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		Build(), nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*organization.Organization, error) {
	var orgResult organization.Organization
	if err := repo.db.Where("id = ?", id).First(&orgResult).Error; err != nil {
		return nil, err
	}
	return &orgResult, nil
}

func (repo *RepositoryPostgres) GetAllByUserId(userId string) ([]organization.Organization, error) {
	var organizations []organization.Organization

	if err := repo.db.
		Table("organizations").
		Joins(
			"INNER JOIN \"organization_users\" ON \"organizations\".id = \"organization_users\"."+
				"organization_id",
		).
		Where("\"organization_users\".user_id = ?", userId).
		Find(&organizations).
		Error; err != nil {
		return nil, err
	}
	return organizations, nil
}
func (repo *RepositoryPostgres) GetByIDAndUserId(
	orgId string,
	userId string,
) (*organization.Organization, error) {
	var orgResult organization.Organization
	if err := repo.db.
		Table("organizations").
		Joins(
			"INNER JOIN \"organization_users\" ou ON \"organizations\"."+
				"id = ou."+"organization_id",
		).
		Where(" ou.user_id = ?", userId).
		Where("\"organizations\".id = ?", orgId).
		Where("ou.is_active = ?", true).
		First(&orgResult).
		Error; err != nil {
		return nil, err
	}
	return &orgResult, nil
}

func (repo *RepositoryPostgres) Create(
	organization *organization.Organization,
	tx *gorm.DB,
) (*organization.Organization, error) {
	if tx == nil {
		tx = repo.db
	}
	if err := tx.Create(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
func (repo *RepositoryPostgres) Update(organization *organization.Organization) (
	*organization.Organization,
	error,
) {
	if err := repo.db.Save(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
func (repo *RepositoryPostgres) Delete(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&organization.Organization{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}
