package organization

import (
	"github.com/kimosapp/poc/internal/core/errors"
	organizationConstants "github.com/kimosapp/poc/internal/core/model/constants/organization"
	permissionConstants "github.com/kimosapp/poc/internal/core/model/constants/permissions"
	"github.com/kimosapp/poc/internal/core/model/entity/organization"
	request "github.com/kimosapp/poc/internal/core/model/request/organizations"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	teamRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/team"
	teamMemberRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/team-member"
	userOrganizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/user-organization"
)

type CreateTeamUseCase struct {
	userOrganizationRepo          userOrganizationRepository.Repository
	teamRepo                      teamRepository.Repository
	teamMemberRepository          teamMemberRepository.Repository
	checkUserHasPermissionUseCase *CheckUserHasPermissionsToMakeAction
	logger                        logging.Logger
}

func NewCreateTeamUseCase(
	userOrganizationRepo userOrganizationRepository.Repository,
	teamRepo teamRepository.Repository,
	teamMemberRepository teamMemberRepository.Repository,
	checkUserHasPermissionUseCase *CheckUserHasPermissionsToMakeAction,
	logger logging.Logger,
) *CreateTeamUseCase {
	return &CreateTeamUseCase{
		userOrganizationRepo:          userOrganizationRepo,
		checkUserHasPermissionUseCase: checkUserHasPermissionUseCase,
		teamRepo:                      teamRepo,
		teamMemberRepository:          teamMemberRepository,
		logger:                        logger,
	}
}

func (uc *CreateTeamUseCase) Handler(
	userId, organizationId string,
	request *request.CreateTeamRequest,
) (*organization.Team, *errors.AppError) {
	tx := uc.teamRepo.BeginTransaction()
	defer tx.Rollback()
	if !uc.checkUserHasPermissionUseCase.Handler(
		userId,
		organizationId,
		[]string{permissionConstants.PERMISSION_CREATE_TEAM},
	) {
		return nil, errors.NewForbiddenError(
			"Error creating team in the organization",
			"The user don't have access to create a team",
			errors.ErrorUserCantCreateTeamIntoOrganization,
		).AppError
	}
	teams, err := uc.teamRepo.GetByNameOrSlugAndOrgId(request.Name, request.Slug, organizationId)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error trying to get organizations and slugs for the company",
			"",
			errors.ErrorTryingToGetTeamsByNameAndSlug,
		).AppError
	}
	if len(teams) > 0 {
		return nil, errors.NewConflictError(
			"The organization has a team with the same name or slug",
			"",
			errors.ErrorConflictTeamExistWithSameNameOrSlug,
		).AppError
	}
	team, err := uc.teamRepo.Create(
		&organization.Team{
			Name:           request.Name,
			Slug:           request.Slug,
			About:          request.About,
			OrganizationID: organizationId,
		}, tx,
	)
	if err != nil {
		uc.logger.Error("Error creating the team", err)
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating team into organization",
			"Error at the moment to create the organization",
			errors.ErrorCreatingTeam,
		).AppError
	}
	// This may can be another case
	_, err = uc.teamMemberRepository.Create(
		&organization.TeamMember{
			TeamID:   team.ID,
			UserID:   userId,
			RoleID:   string(organizationConstants.ADMIN_TEAM),
			IsActive: true,
			Status:   "active",
		}, tx,
	)
	if err != nil {
		uc.logger.Error("Error creating the team member", err)
		tx.Rollback()
		return nil, errors.NewInternalServerError(
			"Error creating team member into team",
			"Error at the moment to create the organization",
			errors.ErrorCreatingTeam,
		).AppError
	}
	tx.Commit()
	return team, nil
}
