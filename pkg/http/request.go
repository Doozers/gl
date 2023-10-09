package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Header struct {
	Key   string
	Value string
}

type Method string

var (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

// GetFromJsonReq is a generic function that sends a request to an external API and returns the value of a field in the response
// field is the name of the field in the response that you want to retrieve, you can omit it if you want the whole response
func GetFromJsonReq[T any](url string, method Method, payload string, headers []Header, field string) (T, error) {
	// default generic value
	var v T

	externalReq, err := http.NewRequest(string(method), url, strings.NewReader(payload))
	if err != nil {
		return v, fmt.Errorf("failed to create request: %w", err)
	}

	for _, header := range headers {
		externalReq.Header.Add(header.Key, header.Value)
	}

	res, err := http.DefaultClient.Do(externalReq)
	if err != nil {
		return v, fmt.Errorf("failed to send request: %w", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return v, fmt.Errorf("failed to read body: %w", err)
	}

	if len(field) == 0 {
		err = json.Unmarshal(body, &v)
		if err != nil {
			return v, fmt.Errorf("failed to unmarshal body: %w", err)
		}
		return v, nil
	}

	var jsonRes map[string]interface{}
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		return v, fmt.Errorf("failed to unmarshal body: %w", err)
	}
	v, ok := jsonRes[field].(T)
	if !ok {
		return v, fmt.Errorf("failed to cast response to type %T", v)
	}
	return v, nil
}
