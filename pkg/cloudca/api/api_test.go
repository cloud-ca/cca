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
	ccaClient := CcaApiClient{server.URL, "api-key", httpClient}

	expectedResp := CcaResponse{
		TaskId:     "test_task_id",
		TaskStatus: "test_task_status",
		Data:       []byte(`{"key":"value"}`),
		MetaData:   map[string]interface{}{"meta_key": "meta_value"},
		StatusCode: 200,
	}

	//when
	resp, _ := ccaClient.Do(CcaRequest{Method: "GET", Endpoint: "/fooo"})

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
	ccaClient := CcaApiClient{server.URL, "api-key", httpClient}

	expectedResp := CcaResponse{
		Errors:     []CcaError{{ErrorCode: "FOO_ERROR", Message: "message1"}, {ErrorCode: "BAR_ERROR", Message: "message2"}},
		StatusCode: 400,
	}

	//when
	resp, _ := ccaClient.Do(CcaRequest{Method: "GET", Endpoint: "/fooo"})

	//then
	assert.Equal(t, expectedResp, *resp)
}
