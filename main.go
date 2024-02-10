package main

import (
	"github.com/Yota-K/aws-cost-notification/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
