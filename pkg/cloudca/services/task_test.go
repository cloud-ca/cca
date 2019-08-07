package services

import (
	"github.com/cloud-ca/cca/pkg/cloudca/api"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/api_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TEST_TASK_ID = "test_task_id"
)

func TestGetTaskReturnTaskIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCcaClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockCcaClient,
	}

	expectedTask := Task{
		Id:      TEST_TASK_ID,
		Status:  "SUCCESS",
		Created: "2015-07-07",
		Result:  []byte(`{"key": "value"}`),
	}

	mockCcaClient.EXPECT().Do(api.CcaRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}).Return(&api.CcaResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"SUCCESS", "created":"2015-07-07", "result":{"key": "value"}}`),
	}, nil)

	//when
	task, _ := taskService.Get(TEST_TASK_ID)

	//then
	assert.Equal(t, expectedTask, *task)
}

func TestGetTaskReturnErrorIfHasCcaErrors(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCcaClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockCcaClient,
	}

	ccaResponse := api.CcaResponse{
		StatusCode: 400,
		Errors:     []api.CcaError{{}},
	}
	mockCcaClient.EXPECT().Do(api.CcaRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}).Return(&ccaResponse, nil)

	//when
	task, err := taskService.Get(TEST_TASK_ID)

	//then
	assert.Nil(t, task)
	assert.Equal(t, api.CcaErrorResponse(ccaResponse), err)
}

func TestGetTaskReturnErrorIfHasUnexpectedErrors(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCcaClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockCcaClient,
	}

	mockError := mocks.MockError{"some_get_task_error"}

	mockCcaClient.EXPECT().Do(api.CcaRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}).Return(nil, mockError)

	//when
	task, err := taskService.Get(TEST_TASK_ID)

	//then
	assert.Nil(t, task)
	assert.Equal(t, mockError, err)
}

func TestPollingReturnTaskResultOnSuccessfulComplete(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCcaClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockCcaClient,
	}

	request := api.CcaRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}

	expectedResult := []byte(`{"foo":"bar"}`)

	pendingResponse := &api.CcaResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"PENDING", "created":"2015-07-07"}`),
	}
	successResponse := &api.CcaResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"SUCCESS", "created":"2015-07-07", "result":` + string(expectedResult) + `}`),
	}
	gomock.InOrder(
		mockCcaClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockCcaClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockCcaClient.EXPECT().Do(request).Return(successResponse, nil),
	)

	//when
	result, _ := taskService.Poll(TEST_TASK_ID, 10)

	//then
	if assert.NotNil(t, result) {
		assert.Equal(t, expectedResult, result)
	}
}

func TestPollingGetErrorOnTaskFailure(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCcaClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockCcaClient,
	}

	request := api.CcaRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}

	expectedResult := []byte(`{"foo":"bar"}`)

	pendingResponse := &api.CcaResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"PENDING", "created":"2015-07-07"}`),
	}
	failedResponse := &api.CcaResponse{
		StatusCode: 400,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"FAILED", "created":"2015-07-07", "result":` + string(expectedResult) + `}`),
	}
	gomock.InOrder(
		mockCcaClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockCcaClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockCcaClient.EXPECT().Do(request).Return(failedResponse, nil),
	)

	//when
	_, err := taskService.Poll(TEST_TASK_ID, 10)

	//then
	assert.NotNil(t, err)

}
