package users

import (
	"github.com/kimosapp/poc/internal/core/errors"
	"github.com/kimosapp/poc/internal/core/model/commons/is_valid"
	"github.com/kimosapp/poc/internal/core/model/entity/users"
	users2 "github.com/kimosapp/poc/internal/core/model/request/users"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	userR "github.com/kimosapp/poc/internal/core/ports/repository/users"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CreateUserUseCase struct {
	userRepository userR.UserRepository
	logger         logging.Logger
}

func NewCreateUserUseCase(
	ur userR.UserRepository,
	logger logging.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: ur, logger: logger}
}

func (p *CreateUserUseCase) Handler(req *users2.SignUpRequest) (
	*users.User,
	*errors.AppError,
) {
	appError := validateSignUpRequest(req)
	if appError != nil {
		return nil, appError
	}
	user, err := p.userRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting users by email",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	if user != nil {
		p.logger.Error("User already exists", "email", req.Email)
		return nil, errors.NewBadRequestError(
			"User already exists",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	hashedPassword, err := hashAndSalt(req.Password)
	if err != nil {
		p.logger.Error("Error hashing password", "errors", err.Error())
		return nil, errors.NewInternalServerError(
			"Error creating the users",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	user = &users.User{
		Email:                      req.Email,
		Hash:                       hashedPassword,
		BadLoginAttempts:           0,
		IsLocked:                   false,
		AcceptTermsAndConditions:   req.AcceptTermsAndConditions,
		AcceptTermsAndConditionsAt: time.Now(),
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}
	createUserResult, err := p.userRepository.Create(user)
	if err != nil {
		p.logger.Error("Error creating users", "errors", err.Error())
		return nil, errors.NewInternalServerError(
			"Error creating users",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	return createUserResult, nil
}

func hashAndSalt(pwd string) (string, error) {
	bytePassword := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func validateSignUpRequest(signUpRequest *users2.SignUpRequest) *errors.AppError {
	if !signUpRequest.AcceptTermsAndConditions {
		return errors.NewBadRequestError(
			"User must accept terms and conditions",
			"",
			errors.ErrorUserNotAcceptTermsAndConditions,
		).AppError
	}
	if !is_valid.IsValidEmail(signUpRequest.Email) {
		return errors.NewBadRequestError(
			"Invalid email",
			"",
			errors.ErrorInvalidEmail,
		).AppError
	}
	if !is_valid.IsValidPassword(signUpRequest.Password) {
		return errors.NewBadRequestError(
			"Invalid password",
			"",
			errors.ErrorPasswordDoesntHaveTheRequestedFormat,
		).AppError
	}
	if signUpRequest.Password != signUpRequest.ConfirmPassword {
		return errors.NewBadRequestError(
			"Password and confirm password don't match",
			"",
			errors.ErrorPasswordDoesntMatch,
		).AppError
	}
	return nil
}
