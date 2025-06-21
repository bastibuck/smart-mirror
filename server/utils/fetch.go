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
	Headers  map[string]string
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

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	resp, err := client.Do(httpReq)
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
