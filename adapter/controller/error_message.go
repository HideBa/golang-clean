package controller

import (
	"fmt"

	"gopkg.in/validator.v2"
)

var validateErrorMessages = map[error]string{
	validator.ErrUnsupported: "%s is invalid value",
	validator.ErrZeroValue:   "please input %s",
	validator.ErrLen:         "the length of string is invalid",
	validator.ErrMax:         "string exceeds the max length",
	ErrRequired:              "Please input %s",
	ErrEmail:                 "%s is invalid email type",
	ErrUint:                  "Please use the number more than 0 instead of %s",
	ErrUniq:                  "%s is already registered",
}

var displayNames = map[string]string{
	"user_id":    "ユーザーID",
	"user_name":  "ユーザー名",
	"article_id": "記事ID",
	"email":      "メールアドレス",
	"content":    "本文",
}

func ConvertErrorsToMessage(errs map[string]error) map[string]string {
	messages := map[string]string{}

	for argName, err := range errs {
		display := displayNames[argName]
		if display == "" {

		}
		message := validateErrorMessages[err]
		if message == "" {
			messages[argName] = err.Error()
		} else {
			messages[argName] = fmt.Sprintf(message)
		}
	}
	return messages
}
