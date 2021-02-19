package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/glog"
)

type Response400Body struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func commonHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}
}

func Response400(errs map[string]error) events.APIGatewayProxyResponse {
	glog.Warningf("%+v", errs)
	res := &Response400Body{
		Message: "Please check the value you input",
		Errors:  ConvertErrorsToMessage(errs),
	}

	b, err := json.Marshal(res)
	if err != nil {
		return Response500(err)
	}
}

func Response500(err error) events.APIGatewayProxyResponse {
	glog.Errorf("%+v\n", err)
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    commonHeaders(),
		Body:       `{"message": "internal server error"}`,
	}
}
