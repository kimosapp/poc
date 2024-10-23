package organization

import (
	"github.com/kimosapp/poc/internal/core/errors"
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	users "github.com/kimosapp/poc/internal/core/model/entity/users"
	request "github.com/kimosapp/poc/internal/core/model/request/organizations"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	organizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations"
	roleRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/role"
	userOrganizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/user-organization"
	userR "github.com/kimosapp/poc/internal/core/ports/repository/users"
	notificationsService "github.com/kimosapp/poc/internal/core/ports/service/notification"
	"github.com/kimosapp/poc/internal/core/utils"
	"golang.org/x/crypto/bcrypt"
)

// TODO update errors
type CreateOrganizationUseCase struct {
	organizationRepository organizationRepository.Repository
	userOrganizationRepo   userOrganizationRepository.Repository
	notificationService    notificationsService.Service
	roleRepo               roleRepository.Repository
	userRepo               userR.Repository
	logger                 logging.Logger
}

func NewCreateOrganizationUseCase(
	organizationRepository organizationRepository.Repository,
	userOrganizationRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	userRepo userR.Repository,
	notificationService notificationsService.Service,
	logger logging.Logger,
) *CreateOrganizationUseCase {
	return &CreateOrganizationUseCase{
		organizationRepository: organizationRepository,
		userOrganizationRepo:   userOrganizationRepo,
		roleRepo:               roleRepo,
		userRepo:               userRepo,
		notificationService:    notificationService,
		logger:                 logger,
	}
}

func (cu CreateOrganizationUseCase) Handler(
	request *request.CreateOrganizationRequest,
) (*organization.Organization, *errors.AppError) {
	persistedOrg, err := cu.organizationRepository.GetByBillingEmail(request.BillingEmail)
	if err != nil {
		cu.logger.Error("Error getting organization by billing email\n", err)
		return nil, errors.NewInternalServerError(
			"Error creating the organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	if persistedOrg.ID != "" {
		cu.logger.Error(
			"Error creating the organization: The billing email has associated"+
				" another organization\n", err,
		)
		return nil, errors.NewBadRequestError(
			"The email "+request.BillingEmail+" has created another organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	tx := cu.organizationRepository.BeginTransaction()
	defer tx.Rollback()
	organizationResult, err := cu.organizationRepository.Create(
		&organization.Organization{
			Name:         request.OrganizationName,
			BillingEmail: request.BillingEmail,
			Slug:         utils.CreateSlug(request.OrganizationName),
		},
		tx,
	)
	if err != nil {
		cu.logger.Error(
			"Error creating the organization: \n", err,
		)
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating the organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	password, err := hashAndSalt(request.Password)
	if err != nil {
		cu.logger.Error(
			"Error creating the hash for the user: \n", err,
		)
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating the organization",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	userCreationResult, err := cu.userRepo.Create(
		&users.User{
			Email:          request.BillingEmail,
			FirstName:      request.FirstName,
			LastName:       request.LastName,
			Hash:           password,
			OrganizationID: organizationResult.ID,
		},
	)
	if err != nil {
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating User",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	organizationResult.CreatedBy = userCreationResult.ID
	organizationResult, err = cu.organizationRepository.Update(organizationResult, tx)
	if err != nil {
		tx.Rollback()
		return organizationResult, errors.NewInternalServerError(
			"Error updating organization",
			"",
			errors.ErrorCreatingOrganization,
		).AppError
	}
	cu.notificationService.SendCreateOrganizationEmail(request.BillingEmail)
	tx.Commit()
	return organizationResult, nil
}

func hashAndSalt(pwd string) (string, error) {
	bytePassword := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
