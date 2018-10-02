package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("main::")
	lambda.Start(Handler)
}

// func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	fmt.Println("Received body: ", request.Body)
// 	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
// }

func Handler(ctx context.Context, name MyEvent) (string, error) {
	fmt.Println("handler:: ")
	return "hi", nil
}
