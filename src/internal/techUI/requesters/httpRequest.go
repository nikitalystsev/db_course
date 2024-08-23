package requesters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HTTPRequest struct {
	Method      string
	URL         string
	Headers     map[string]string
	Body        interface{}
	QueryParams map[string]string
	Timeout     time.Duration
}

type HTTPResponse struct {
	Status     string
	StatusCode int
	Headers    map[string][]string
	Body       []byte
}

func SendRequest(req HTTPRequest) (*HTTPResponse, error) {
	parsedURL, err := url.Parse(req.URL)
	if err != nil {
		return nil, err
	}

	if len(req.QueryParams) > 0 {
		q := parsedURL.Query()
		for key, value := range req.QueryParams {
			q.Add(key, value)
		}
		parsedURL.RawQuery = q.Encode()
	}

	var body io.Reader
	if req.Body != nil {
		var jsonBody []byte
		if jsonBody, err = json.Marshal(req.Body); err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBody)
	}

	httpReq, err := http.NewRequest(req.Method, parsedURL.String(), body)
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		httpReq.Header.Add(key, value)
	}

	client := &http.Client{
		Timeout: req.Timeout,
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("error closing body")
		}
	}(resp.Body)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	httpResp := &HTTPResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       responseBody,
	}

	return httpResp, nil
}
