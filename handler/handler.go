package handler

import (
	"context"
	"github.com/Yota-K/aws-cost-notification/notification"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

type Response events.APIGatewayProxyResponse

func errorResponse(err error) (Response, error) {
	return Response{
		StatusCode: 500,
		Body:       "Error",
	}, err
}

func HandleRequest(ctx context.Context) (Response, error) {
	api, err := notification.NewApi()
	if err != nil {
		log.Fatal(err)
		errorResponse(err)
	}

	result, err := notification.GetCostMonthly(api)
	if err != nil {
		log.Fatal(err)
		errorResponse(err)
	}

	month := result.ResultsByTime[0].TimePeriod
	cost := result.ResultsByTime[0].Total["UnblendedCost"]

	amount := *cost.Amount
	currentJpy := notification.GetCurrentJpy()
	jpy := notification.ConvertUsDollarToJpy(amount, currentJpy)

	message := `
開始月: ` + *month.Start + `
終了月: ` + *month.End + `

今月のAWS利用料金は、` + jpy + `円です。`

	notification.SendMessage(message)

	return Response{
		StatusCode: 200,
		Body:       "success",
	}, nil
}
