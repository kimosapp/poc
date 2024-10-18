package errors

type ErrorCode string

const (
	//Users
	ErrorUserNotAcceptTermsAndConditions      ErrorCode = "0000001"
	ErrorInvalidEmail                         ErrorCode = "0000002"
	ErrorPasswordDoesntHaveTheRequestedFormat ErrorCode = "0000003"
	ErrorPasswordDoesntMatch                  ErrorCode = "0000004"
	ErrorCreatingUser                         ErrorCode = "0000005"
	ErrorAuthenticatingUser                   ErrorCode = "0000006"
	ErrorUserAuthenticatedNotFound            ErrorCode = "0000007"
	ErrorUserEmailAlreadyExists               ErrorCode = "0000008"

	// Organizations
	ErrorCreatingOrganization                         ErrorCode = "0000026"
	ErrorUserCantAddUsersToOrganization               ErrorCode = "0000009"
	ErrorUserCantRemoveUsersFromOrganization          ErrorCode = "00000020"
	ErrorUserCantCreateTeamIntoOrganization           ErrorCode = "00000021"
	ErrorTryingToGetTeamsByNameAndSlug                ErrorCode = "00000022"
	ErrorConflictTeamExistWithSameNameOrSlug          ErrorCode = "00000023"
	ErrorCreatingTeam                                 ErrorCode = "00000024"
	ErrorUserDontHavePrivilegesToAddTeamMembersToTeam ErrorCode = "00000025"
)
