package organization

import (
	userOrganizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/user-organization"
	userRepository "github.com/kimosapp/poc/internal/core/ports/repository/users"
)

type GetTeamMembersPaginatedUseCase struct {
	userOrganizationRepo *userOrganizationRepository.Repository
	userRepo             *userRepository.Repository
}

func NewGetTeamMembersPaginatedUseCase(
	userOrganizationRepo *userOrganizationRepository.Repository,
	userRepo *userRepository.Repository,
) *GetTeamMembersPaginatedUseCase {
	return &GetTeamMembersPaginatedUseCase{
		userOrganizationRepo: userOrganizationRepo,
		userRepo:             userRepo,
	}

}
