package notification

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	BASE_LINE_URL = "https://notify-api.line.me/api/notify"
)

func SendMessage(message string) {
	token := os.Getenv("LINE_ACCESS_TOKEN")

	u, err := url.ParseRequestURI(BASE_LINE_URL)
	if err != nil {
		log.Fatal(err)
	}

	c := http.Client{}

	v := url.Values{}
	v.Set("message", message)
	body := strings.NewReader(v.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("status: %s", res.Status)
}
