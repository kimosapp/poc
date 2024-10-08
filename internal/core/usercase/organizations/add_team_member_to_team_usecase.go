package organization

import (
	"github.com/kimosapp/poc/internal/core/errors"
	organizationRequest "github.com/kimosapp/poc/internal/core/model/request/organizations"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	teamRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/team"
	teamMemberRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/team-member"
	userOrganizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/user-organization"
)

type AddTeamMembersUseCase struct {
	userOrganizationRepo                  userOrganizationRepository.Repository
	teamRepo                              teamRepository.Repository
	teamMemberRepo                        teamMemberRepository.Repository
	checkIfUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction
	logger                                logging.Logger
}

func NewAddTeamMembersUseCase(
	userOrganizationRepo userOrganizationRepository.Repository,
	teamRepo teamRepository.Repository,
	teamMemberRepo teamMemberRepository.Repository,
	checkIfUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction,
	logger logging.Logger,
) *AddTeamMembersUseCase {
	return &AddTeamMembersUseCase{
		userOrganizationRepo:                  userOrganizationRepo,
		teamRepo:                              teamRepo,
		teamMemberRepo:                        teamMemberRepo,
		checkIfUserHasPermissionsToMakeAction: checkIfUserHasPermissionsToMakeAction,
		logger:                                logger,
	}
}

func (u *AddTeamMembersUseCase) Handler(
	authenticatedUserId string,
	orgId string,
	request *organizationRequest.AddTeamMembersRequest,
) *errors.AppError {
	tx := u.teamRepo.BeginTransaction()
	defer tx.Rollback()
	if !u.checkIfUserHasPermissionsToMakeAction.Handler(
		authenticatedUserId,
		orgId,
		[]string{"ADD_TEAM_MEMBERS"},
	) {
		return errors.NewForbiddenError(
			"The user don't have the privileges to do this operation",
			"The user don't have the privileges to do this operation if the error persist, contact with your administrator or contact us",
			errors.ErrorUserDontHavePrivilegesToAddTeamMembersToTeam,
		).AppError
	}
	// I should be able to invite a new member form here  (To org too)
	return nil
}
