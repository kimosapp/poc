package errors

type ErrorCode string

const (
	ErrorUserNotAcceptTermsAndConditions      ErrorCode = "0000001"
	ErrorInvalidEmail                         ErrorCode = "0000002"
	ErrorPasswordDoesntHaveTheRequestedFormat ErrorCode = "0000003"
	ErrorPasswordDoesntMatch                  ErrorCode = "0000004"
	ErrorCreatingUser                         ErrorCode = "0000005"
	ErrorAuthenticatingUser                   ErrorCode = "0000006"
	ErrorUserAuthenticatedNotFound            ErrorCode = "0000007"
	ErrorUserEmailAlreadyExists               ErrorCode = "0000008"
)
