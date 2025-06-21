package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RelaxedHttpRequestOptions struct {
	URL      string
	Method   string
	Response interface{}
}

func RelaxedHttpRequest(req RelaxedHttpRequestOptions) error {
	if req.Method == "" {
		req.Method = "GET" // Default to GET if no method is specified
	}

	client := &http.Client{}
	httpReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(req.Response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
