package interactor

import (
	"github.com/HideBa/golang-clean/domain"
	"github.com/HideBa/golang-clean/usecase"
	"github.com/pkg/errors"
)

var ErrUniqueEmail = errors.New("email is not unique")

type UserCreater struct {
	UserRepository  domain.UserRepository
	UniqueValidator *domain.UserEmailUniqueValidator
}

func NewUserCreater(repository domain.UserRepository, validator *domain.UserEmailUniqueValidator) *UserCreater {
	return &UserCreater{
		UserRepository:  repository,
		UniqueValidator: validator,
	}
}

func (u *UserCreater) Execute(req *usecase.CreateUserRequest) (*usecase.CreateUserResponse, error) {
	isUnique, err := u.UniqueValidator.IsUniqueEmail(req.ToUserModel())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !isUnique {
		return nil, errors.WithStack(ErrUniqueEmail)
	}

	user, err := u.UserRepository.CreateUser(req.ToUserModel())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &usecase.CreateUserResponse{User: user}, nil
}
