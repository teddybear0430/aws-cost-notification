package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"testing"
)

type MockCostExplorerApi struct {
	output *costexplorer.GetCostAndUsageOutput
	Error  error
}

func (m *MockCostExplorerApi) GetCostAndUsage(input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	return m.output, m.Error
}

func TestGetCostMonthly(t *testing.T) {
	mockData := &MockCostExplorerApi{
		output: &costexplorer.GetCostAndUsageOutput{
			DimensionValueAttributes: []*costexplorer.DimensionValuesWithAttributes{},
			ResultsByTime: []*costexplorer.ResultByTime{
				{
					Estimated: aws.Bool(false),
					Groups:    []*costexplorer.Group{},
					TimePeriod: &costexplorer.DateInterval{
						Start: aws.String("2024-01-01"),
						End:   aws.String("2024-02-01"),
					},
					Total: map[string]*costexplorer.MetricValue{
						"UnblendedCost": {
							Amount: aws.String("1.4915533215"),
							Unit:   aws.String("USD"),
						},
					},
				},
			},
		},
		Error: nil,
	}

	result, err := GetCostMonthly(mockData)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if *result.ResultsByTime[0].TimePeriod.Start != "2024-01-01" {
		t.Errorf("Start: %v", *result.ResultsByTime[0].TimePeriod.Start)
	}
	if *result.ResultsByTime[0].TimePeriod.End != "2024-02-01" {
		t.Errorf("End: %v", *result.ResultsByTime[0].TimePeriod.End)
	}
	if *result.ResultsByTime[0].Total["UnblendedCost"].Amount != "1.4915533215" {
		t.Errorf("Amount: %v", *result.ResultsByTime[0].Total["UnblendedCost"].Amount)
	}
	if *result.ResultsByTime[0].Total["UnblendedCost"].Unit != "USD" {
		t.Errorf("Unit: %v", *result.ResultsByTime[0].Total["UnblendedCost"].Unit)
	}
}
