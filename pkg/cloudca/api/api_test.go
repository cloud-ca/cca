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
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskReturnTaskIfSuccess(t *testing.T) {
	//given
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"taskId": "test_task_id", `+
			`"taskStatus": "test_task_status", `+
			`"data": {"key":"value"}, `+
			`"metadata": {"meta_key":"meta_value"}}`)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	ccaClient := CcaClient{server.URL, "api-key", httpClient}

	expectedResp := Response{
		TaskID:     "test_task_id",
		TaskStatus: "test_task_status",
		Data:       []byte(`{"key":"value"}`),
		MetaData:   map[string]interface{}{"meta_key": "meta_value"},
		StatusCode: 200,
	}

	//when
	resp, _ := ccaClient.Do(Request{Method: "GET", Endpoint: "/fooo"})

	//then
	assert.Equal(t, expectedResp, *resp)
}

func TestGetTaskReturnErrorsIfErrorOccured(t *testing.T) {
	//given
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"errors": [{"errorCode": "FOO_ERROR", "message": "message1"}, {"errorCode": "BAR_ERROR", "message":"message2"}]}`)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	ccaClient := CcaClient{server.URL, "api-key", httpClient}

	expectedResp := Response{
		Errors:     []Error{{ErrorCode: "FOO_ERROR", Message: "message1"}, {ErrorCode: "BAR_ERROR", Message: "message2"}},
		StatusCode: 400,
	}

	//when
	resp, _ := ccaClient.Do(Request{Method: "GET", Endpoint: "/fooo"})

	//then
	assert.Equal(t, expectedResp, *resp)
}
