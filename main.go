package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type user struct {
	Name   string
	Active bool
	Role   string
}

func handlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processando Request ID %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body Size -> %d.\n", len(request.Body))

	u := user{
		Name:   "User Name",
		Active: true,
		Role:   "Admin",
	}

	responseUserFormat, err := json.Marshal(&u)

	var response events.APIGatewayProxyResponse

	if err != nil {
		response = events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Erro na resposta -> %s", err.Error()),
			StatusCode: http.StatusBadRequest,
		}
	}

	response = events.APIGatewayProxyResponse{
		Body:       string(responseUserFormat),
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func main() {
	lambda.Start(handlerRequest)
}
