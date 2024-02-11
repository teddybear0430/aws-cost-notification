package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"log/slog"
	"time"
)

// AWS APIのモック用にインターフェースとして切り出す
type CostExplorerApi interface {
	GetCostAndUsage(input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error)
}

// AWS APIの実態を返却する関数
func NewApi() (CostExplorerApi, error) {
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("ap-northeast-1"),
		},
	}))

	return costexplorer.New(ses), nil
}

// 月のAWS利用料金を取得する
func GetCostMonthly(api CostExplorerApi) (*costexplorer.GetCostAndUsageOutput, error) {
	// 現在の日付と1ヶ月前の日付を日本標準時で取得する
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)
	oneMonthAgo := now.AddDate(0, -1, 0)
	dateBeforeOneMonth := oneMonthAgo.Format("2006-01-02")
	nowDate := now.Format("2006-01-02")

	slog.Info("開始月", "start", dateBeforeOneMonth)
	slog.Info("終了月", "end", nowDate)

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(dateBeforeOneMonth),
			End:   aws.String(nowDate),
		},
	}

	result, err := api.GetCostAndUsage(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
