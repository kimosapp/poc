package users

import (
	"github.com/kimosapp/poc/internal/core/errors"
	"github.com/kimosapp/poc/internal/core/model/response/users"
	"github.com/kimosapp/poc/internal/core/ports/logging"
	repository "github.com/kimosapp/poc/internal/core/ports/repository/users"
)

type GetUserUseCase struct {
	userRepository repository.Repository
	logger         logging.Logger
}

func NewGetUserUseCase(
	ur repository.Repository,
	logger logging.Logger,
) *GetUserUseCase {
	return &GetUserUseCase{userRepository: ur, logger: logger}
}

func (p *GetUserUseCase) Handler(id string) (
	*users.UserLightDTO,
	*errors.AppError,
) {
	result, err := p.userRepository.GetByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
			"",
			errors.ErrorUserAuthenticatedNotFound,
		).AppError
	}
	if result == nil {
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
			"",
			errors.ErrorUserAuthenticatedNotFound,
		).AppError
	}
	return &users.UserLightDTO{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		LastLogin: result.LastLogin,
		CreatedAt: result.CreatedAt,
	}, nil
}
