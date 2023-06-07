// Package api provides a Kagi API client.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient(tkn string) *Client {
	return &Client{
		http: &http.Client{
			Transport: &roundTripper{tkn: tkn},
		},
	}
}

type FastGPTResponse struct {
	Meta   FastGPTResponseMeta    `json:"meta"`
	Data   FastGPTResponseData    `json:"data"`
	Errors []FastGPTResponseError `json:"error"`
}

type FastGPTResponseMeta struct {
	ID           string `json:"id"`
	Node         string `json:"node"`
	Milliseconds int    `json:"ms"`
}

type FastGPTResponseData struct {
	Output     string                     `json:"output"`
	Tokens     int                        `json:"tokens"`
	References []FastGPTResponseReference `json:"references"`
}

type FastGPTResponseReference struct {
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	Link    string `json:"url"`
}

type FastGPTResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	// There's also "ref", but it was null so I don't know its type.
}

type FastGPTRequest struct {
	Query     string `json:"query"`
	WebSearch bool   `json:"web_search"`
	Cache     bool   `json:"cache"`
}

func (c *Client) QueryFastGPT(query string) (*FastGPTResponse, error) {
	var buf bytes.Buffer
	req := FastGPTRequest{
		Query:     query,
		WebSearch: true,
		Cache:     true,
	}
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}
	resp, err := c.http.Post("https://kagi.com/api/v0/fastgpt", "application/json", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to query FastGPT API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Best-effort read the response body.
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d, body %q", resp.StatusCode, string(body))
	}

	var gptResp FastGPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&gptResp); err != nil {
		return nil, fmt.Errorf("failed to decode FastGPT API response body: %w", err)
	}

	if len(gptResp.Errors) > 0 {
		e := gptResp.Errors[0]
		return nil, fmt.Errorf("received %d error(s) from the API: [%d] %q", len(gptResp.Errors), e.Code, e.Message)
	}

	return &gptResp, nil
}

type roundTripper struct {
	tkn string
}

func (rt *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bot "+rt.tkn)
	return http.DefaultTransport.RoundTrip(r)
}
