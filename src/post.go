package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// fmt.Println("Received body: ", request.Body)
	// sess, _ := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-2")},
	// )
	// svc := translate.New(sess)
	return events.APIGatewayProxyRequest{Body: request.Body, StatusCode: 200}, nil

	//return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
