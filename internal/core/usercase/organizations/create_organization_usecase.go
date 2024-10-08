package organization

import (
	"github.com/kimosapp/poc/internal/core/errors"
	roleConstant "github.com/kimosapp/poc/internal/core/model/constants/roles"
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	request "github.com/kimosapp/poc/internal/core/model/request/organizations"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	organizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations"
	roleRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/role"
	userOrganizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/user-organization"
	userR "github.com/kimosapp/poc/internal/core/ports/repository/users"
	"github.com/kimosapp/poc/internal/core/utils"
	"time"
)

// TODO update errors
type CreateOrganizationUseCase struct {
	organizationRepository organizationRepository.Repository
	userOrganizationRepo   userOrganizationRepository.Repository
	roleRepo               roleRepository.Repository
	userRepo               userR.Repository
	logger                 logging.Logger
}

func NewCreateOrganizationUseCase(
	organizationRepository organizationRepository.Repository,
	userOrganizationRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	userRepo userR.Repository,
	logger logging.Logger,
) *CreateOrganizationUseCase {
	return &CreateOrganizationUseCase{
		organizationRepository: organizationRepository,
		userOrganizationRepo:   userOrganizationRepo,
		roleRepo:               roleRepo,
		userRepo:               userRepo,
		logger:                 logger,
	}
}

func (cu CreateOrganizationUseCase) Handler(
	userId string,
	request *request.CreateOrganizationRequest,
) (*organization.Organization, *errors.AppError) {
	tx := cu.organizationRepository.BeginTransaction()
	defer tx.Rollback()

	organizationResult, err := cu.organizationRepository.Create(
		&organization.Organization{
			Name:         request.Name,
			BillingEmail: request.BillingEmail,
			CreatedBy:    userId,
			Slug:         utils.CreateSlug(request.Name),
		},
		tx,
	)
	if err != nil {
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating the organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	userEntityResult, err := cu.userRepo.GetByID(userId)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error creating user organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	_, err = cu.userOrganizationRepo.Create(
		&organization.UserOrganization{
			OrganizationID: organizationResult.ID,
			UserID:         userEntityResult.ID,
			RoleID:         roleConstant.ORGANIZATION_ADMIN,
			//TODO send email to user with different template to the invite
			InvitedAt: time.Now(),
		}, tx,
	)
	if err != nil {
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating user organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	tx.Commit()
	return organizationResult, nil
}
