package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Valeu string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process request ID %f", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
