package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"log/slog"
	"time"
)

type costClient struct {
	ses *session.Session
	now time.Time
}

func NewCostExplorerClient() *costClient {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("ap-northeast-1"),
		},
	}))

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)

	return &costClient{ses: ses, now: now}
}

// 月のAWS利用料金を取得する
func (c *costClient) GetCostMonthly() (*costexplorer.GetCostAndUsageOutput, error) {
	svc := costexplorer.New(c.ses)

	// 現在の日付と1ヶ月前の日付を日本標準時で取得する
	oneMonthAgo := c.now.AddDate(0, -1, 0)
	dateBeforeOneMonth := oneMonthAgo.Format("2006-01-02")
	nowDate := c.now.Format("2006-01-02")

	slog.Info("開始月", "start", dateBeforeOneMonth)
	slog.Info("終了月", "end", nowDate)

	// TODO: 後で実行月の日付を指定するように変える
	input := &costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2023-12-01"),
			End:   aws.String("2024-01-01"),
			// Start: aws.String(dateBeforeOneMonth),
			// End:   aws.String(nowDate),
		},
	}

	result, err := svc.GetCostAndUsage(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
