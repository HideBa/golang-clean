package usecase

import "github.com/HideBa/golang-clean/domain"

type ICreateUser interface {
	Execute(req *CreateUserRequest) (*CreateUserRequest, error)
}

type CreateUserRequest struct {
	Name  string
	Email string
}

func (u *CreateUserRequest) ToUserModel() *domain.UserModel {
	return domain.NewUserModel(u.Name, u.Email)
}

type CreateUserResponse struct {
	User *domain.UserModel
}

func (u *CreateUserResponse) GetUserID() uint64 {
	return u.User.ID
}
