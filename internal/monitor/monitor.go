package monitor

import (
	"net/http"
	"time"
)

type Result struct {
	URL     string
	Status  string
	Latency time.Duration
}

// CheckService performs the HTTP GET
func CheckService(url string) Result {
	start := time.Now()
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	elapsed := time.Since(start)

	if err != nil || resp.StatusCode != http.StatusOK {
		return Result{URL: url, Status: "DOWN", Latency: elapsed}
	}
	
	defer resp.Body.Close()
	return Result{URL: url, Status: "UP", Latency: elapsed}
}