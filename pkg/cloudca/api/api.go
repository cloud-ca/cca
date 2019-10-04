// Copyright © 2019 cloud.ca Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client Represent the Client interface for interacting with cloud.ca API
type Client interface {
	Do(request Request) (*Response, error)
	GetAPIURL() string
	GetAPIKey() string
}

// CcaClient for interacting with cloud.ca API
type CcaClient struct {
	apiURL     string
	apiKey     string
	httpClient *http.Client
}

// Do Execute the API call to server and returns a Response. cloud.ca errors will
// be returned in the Response body, not in the error return value. The error
// return value is reserved for unexpected errors.
func (c CcaClient) Do(request Request) (*Response, error) {
	var bodyBuffer io.Reader
	if request.Body != nil {
		bodyBuffer = bytes.NewBuffer(request.Body)
	}
	method := request.Method
	if method == "" {
		method = "GET"
	}
	req, err := http.NewRequest(request.Method, c.buildURL(request.Endpoint, request.Options), bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("MC-Api-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		fmt.Printf("%s", err)
	}()
	return NewResponse(resp)
}

// GetAPIKey Return the API key being used by API client
func (c CcaClient) GetAPIKey() string {
	return c.apiKey
}

// GetAPIURL Return the API URL being used by API client
func (c CcaClient) GetAPIURL() string {
	return c.apiURL
}

// NewClient Create a new Client with provided API URL and key
func NewClient(apiURL, apiKey string) Client {
	return CcaClient{
		apiURL:     apiURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// NewInsecureClient Create a new Client with provided API URL and key that accepts insecure connections
func NewInsecureClient(apiURL, apiKey string) Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return CcaClient{
		apiURL:     apiURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Transport: tr},
	}
}

// buildURL Builds a URL by using endpoint and options. Options will be set as query parameters.
func (c CcaClient) buildURL(endpoint string, options map[string]string) string {
	query := url.Values{}
	if options != nil {
		for k, v := range options {
			query.Add(k, v)
		}
	}
	u, _ := url.Parse(c.apiURL + "/" + strings.Trim(endpoint, "/") + "?" + query.Encode())
	return u.String()
}
