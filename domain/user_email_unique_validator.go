package domain

import "github.com/pkg/errors"

type UserEmailUniqueValidator struct {
	Repository UserRepository
}

func NewUserEmailUniqueValidator(repository UserRepository) *UserEmailUniqueValidator {
	return &UserEmailUniqueValidator{Repository: repository}
}

func (u *UserEmailUniqueValidator) IsUniqueEmail(newUser *UserModel) (bool, error) {
	user, err := u.Repository.GetUserByEmail(newUser.Email)
	if err != nil {
		if ErrNotFound.Error() == err.Error() {
			return true, nil
		}
		return false, errors.WithStack(err)
	}

	return user.ID == newUser.ID, nil
}
