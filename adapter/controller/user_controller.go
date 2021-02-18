package controller

import "github.com/aws/aws-lambda-go/events"

type RequestPostUser struct {
	Name  string `json:"user_name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersResponse struct {
	Users []*UserResponse `json:"users"`
}

func PostUsers(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	validator := PostSettingValidator()
	validErr := validator.ValidateBody(request.Body)
	if validErr != nil {
		return Response400(validErr)
	}

	var req RequestPostUser
}

func PostSettingValidator() *Validator {
	return &Validator{
		Settings: []*ValidatorSetting{
			{ArgName: "user_name", ValidateTags: "required"},
			{ArgName: "email", ValidateTags: "required,email"},
		},
	}
}
