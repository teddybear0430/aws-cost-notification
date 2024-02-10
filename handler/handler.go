package handler

import (
	"context"
	"fmt"
	"github.com/Yota-K/aws-cost-notification/notification"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context) (Response, error) {
	c := notification.NewCostExplorerClient()
	result, err := c.GetCostMonthly()

	if err != nil {
		return Response{
			StatusCode: 500,
			Body:       "Error",
		}, err
	}

	fmt.Println(result)

	return Response{
		StatusCode: 200,
		Body:       "success",
	}, nil
}
