package controller

import "gopkg.in/validator.v2"

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
