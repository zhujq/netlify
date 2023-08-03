package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	str := "Headers：\r\n"
	for k, v := range request.Headers {
		str += (k + ":" + v + "\r\n")
	}

	return events.APIGatewayProxyResponse{Body: str, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
