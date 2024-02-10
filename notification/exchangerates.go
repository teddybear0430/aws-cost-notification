package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	BASE_OPENEXCHANGER_URL = "https://openexchangerates.org/api/latest.json?app_id=%s"
)

type OpenExchangeResult struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  json.Number        `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

// 最新の日本円を取得
func GetCurrentJpy() float64 {
	appId := os.Getenv("OPEN_EXCHANGE_RATES_APP_ID")

	endpoint := fmt.Sprintf(BASE_OPENEXCHANGER_URL, appId)
	u, err := url.ParseRequestURI(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	c := http.Client{}
	res, err := c.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	var r OpenExchangeResult
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatal(err)
	}

	return r.Rates["JPY"]
}
