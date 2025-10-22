package controllers

import (
	"down_detective/helpers"
	"down_detective/model"
	"encoding/json"
	"net/http"
	"time"
)

func PingURLsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var URLRequet model.URLRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&URLRequet)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(URLRequet.URLs) == 0 {
		http.Error(w, "No URLs provided", http.StatusBadRequest)
		return
	}

	// default timeout if not provided
	timeout := 5 * time.Second
	if URLRequet.Timeout > 0 {
		timeout = time.Duration(URLRequet.Timeout) * time.Second
	}

	// map to hold results of goroutines
	results := make(chan map[string]string)
	defer close(results)

	// running goroutines to ping URLs
	for _, url := range URLRequet.URLs {
		go helpers.PingURL(url, timeout, results)
	}

	// now store results and return
	var response []map[string]string
	upCount := 0

	for i := 0; i < len(URLRequet.URLs); i++ {
		res := <-results
		response = append(response, res)
		if res["status"] == "up" {
			upCount++
		}
	}

	stats := map[string]interface{}{
		"total_urls": len(URLRequet.URLs),
		"up_urls":    upCount,
		"down_urls":  len(URLRequet.URLs) - upCount,
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"results": response,
		"stats":   stats,
	})

}
