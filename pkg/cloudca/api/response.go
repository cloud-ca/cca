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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//nolint
const (
	OK              = 200
	MultipleChoices = 300
	BadRequest      = 400
	NotFound        = 404
)

// Error API error
type Error struct {
	ErrorCode string                 `json:"errorCode"`
	Message   string                 `json:"message"`
	Context   map[string]interface{} `json:"context"`
}

// Response API Response
type Response struct {
	TaskID     string
	TaskStatus string
	StatusCode int
	Data       []byte
	Errors     []Error
	MetaData   map[string]interface{}
}

// NewResponse returns new response instance based on actual HTTP response
func NewResponse(r *http.Response) (*Response, error) {
	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	response := Response{}
	response.StatusCode = r.StatusCode
	responseMap := map[string]*json.RawMessage{}
	err = json.Unmarshal(respBody, &responseMap)
	if err != nil {
		return nil, err
	}

	if val, ok := responseMap["taskId"]; ok {
		err = json.Unmarshal(*val, &response.TaskID)
		if err != nil {
			return nil, err
		}
	}

	if val, ok := responseMap["taskStatus"]; ok {
		err = json.Unmarshal(*val, &response.TaskStatus)
		if err != nil {
			return nil, err
		}
	}

	if val, ok := responseMap["data"]; ok {
		response.Data = []byte(*val)
	}

	if val, ok := responseMap["metadata"]; ok {
		metadata := map[string]interface{}{}
		err = json.Unmarshal(*val, &metadata)
		if err != nil {
			return nil, err
		}
		response.MetaData = metadata
	}

	if val, ok := responseMap["errors"]; ok {
		errors := []Error{}
		err = json.Unmarshal(*val, &errors)
		if err != nil {
			return nil, err
		}
		response.Errors = errors
	} else if !isInOKRange(r.StatusCode) {
		return nil, fmt.Errorf("Unexpected. Received status " + r.Status + " but no errors in response body")
	}

	return &response, nil
}

// IsError returns true if API response has errors
func (r Response) IsError() bool {
	return !isInOKRange(r.StatusCode)
}

// ErrorResponse API Response with errors
type ErrorResponse Response

// Error
func (r ErrorResponse) Error() string {
	var errorStr = "[ERROR] Received HTTP status code " + strconv.Itoa(r.StatusCode) + "\n"
	for _, e := range r.Errors {
		context, _ := json.Marshal(e.Context)
		errorStr += "[ERROR] Error Code: " + e.ErrorCode + ", Message: " + e.Message + ", Context: " + string(context) + "\n"
	}
	return errorStr
}

func isInOKRange(statusCode int) bool {
	return statusCode >= OK && statusCode < MultipleChoices
}
