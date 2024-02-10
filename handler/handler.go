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

	// TODO: デバッグ用
	fmt.Println(result)

	month := result.ResultsByTime[0].TimePeriod
	cost := result.ResultsByTime[0].Total["UnblendedCost"]

	amount := *cost.Amount
	currentJpy := notification.GetCurrentJpy()
	jpy := notification.ConvertUsDollarToJpy(amount, currentJpy)

	message := `
	開始月: ` + *month.Start + `
	終了月: ` + *month.End + `

	今月のAWS利用料金は、` + jpy + `円です。
	`

	notification.SendMessage(message)

	return Response{
		StatusCode: 200,
		Body:       "success",
	}, nil
}
