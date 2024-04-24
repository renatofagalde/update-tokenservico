package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s! \n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s%s", greeting, "🚀🚀 😎 log pipeline 03"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
