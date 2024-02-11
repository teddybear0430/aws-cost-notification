package notification

import (
	"testing"
)

func TestConvertUsDollarToJpy(t *testing.T) {
	mockData := OpenExchangeResult{
		Rates: map[string]float64{
			"JPY": 149.30498545,
		},
	}
	amount := "3.5442683294"
	currentJpy := mockData.Rates["JPY"]
	expected := "529"
	actual := ConvertUsDollarToJpy(amount, currentJpy)

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
