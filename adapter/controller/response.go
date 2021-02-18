package controller

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/glog"
)

type Response400Body struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func Response400(err map[string]string) events.APIGatewayProxyResponse {
	glog.Warningf("%+v", errs)
	res := &Response400Body{
		Message: "Please check the value you input",
		Errors:  ConvertErrorsToMessage(errs),
	}
}
