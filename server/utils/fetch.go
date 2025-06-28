package utils

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

type RelaxedHttpRequestOptions struct {
	URL      string
	Method   string
	Response interface{} // The response should be a pointer to a struct where the JSON response will be decoded into
	Headers  map[string]string
	Delay    RelaxedHttpRequestDelay
	Timeout  time.Duration
	Retries  int
}

type RelaxedHttpRequestDelay struct {
	Variance int // percent
	Average  int // milliseconds
} // Delay before making the request to avoid rate limiting. It can be modified with a variance to simulate more realistic behavior

func RelaxedHttpRequest(req RelaxedHttpRequestOptions) error {
	if req.Method == "" {
		req.Method = "GET" // Default to GET if no method is specified
	}

	client := &http.Client{
		Timeout: req.Timeout,
	}
	httpReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	if req.Delay.Average > 0 {
		variance := float64(req.Delay.Average) * float64(req.Delay.Variance) / 100.0

		min := float64(req.Delay.Average) - variance
		max := float64(req.Delay.Average) + variance

		randomDelay := min + rand.Float64()*(max-min)

		time.Sleep(time.Duration(randomDelay) * time.Millisecond)
	}

	retries := 0
	if req.Retries > 0 {
		retries = req.Retries
	}
	maxTries := retries + 1 // at least one try

	var resp *http.Response

	try := 0
	for try < maxTries {
		if try > 0 {
			fmt.Printf("[FETCH] Retrying failed request to %s\n", req.URL)
		}

		time.Sleep(time.Duration(try*try*500) * time.Millisecond) // incremental backoff

		resp, err = client.Do(httpReq)

		if err == nil {
			break
		}

		try++
	}

	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
		}

		return fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(req.Response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
