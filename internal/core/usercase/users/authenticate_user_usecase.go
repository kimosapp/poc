package users

import (
	"github.com/kimosapp/poc/internal/core/auth"
	"github.com/kimosapp/poc/internal/core/errors"
	"github.com/kimosapp/poc/internal/core/model/request/users"
	"github.com/kimosapp/poc/internal/core/model/response"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	userR "github.com/kimosapp/poc/internal/core/ports/repository/users"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticateUserUseCase struct {
	userRepository userR.UserRepository
	logger         logging.Logger
}

func NewAuthenticateUserUseCase(
	ur userR.UserRepository,
	logger logging.Logger,
) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{userRepository: ur, logger: logger}
}

func (p *AuthenticateUserUseCase) Handler(request users.LoginRequest) (
	*response.BaseAuthenticationResponse,
	*errors.AppError,
) {
	result, err := p.userRepository.GetByEmail(request.Email)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting users by email",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if result == nil {
		p.logger.Error("User doesn't exist", "email", request.Email)
		return nil, errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if result.IsLocked || result.BadLoginAttempts >= 5 {
		p.logger.Error("User is locked", "email", request.Email)
		return nil, errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if !comparePasswords(result.Hash, request.Password) {
		return nil, errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	expirationTime := time.Now().Add(60 * time.Minute)
	jwt, err := auth.GenerateJWT(result.ID, result.Email, expirationTime)
	if err != nil {
		p.logger.Error("Error generating JWT", "errors", err.Error())
		return nil, errors.NewInternalServerError(
			"Error generating JWT",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return &response.BaseAuthenticationResponse{
		AccessToken: jwt,
		ExpiresIn:   int(expirationTime.Unix()),
		TokenType:   "Bearer",
	}, nil

}

func comparePasswords(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
