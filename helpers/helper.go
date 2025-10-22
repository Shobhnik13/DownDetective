package helpers

import (
	"net/http"
	"time"
)

func PingURL(url string, timeout time.Duration, results chan<- map[string]string) {
	// implement ping logic here and send result to results channel
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	status := ""

	if err != nil {
		status = "down"
	} else {
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			status = "up"
		} else {
			status = "down"
		}
		resp.Body.Close()
	}
	results <- map[string]string{
		"url":    url,
		"status": status,
	}
	
}
