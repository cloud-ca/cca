package configuration

import (
	"encoding/json"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/configuration_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TEST_SERVICE_CONNECTION_ID           = "connection_id"
	TEST_SERVICE_CONNECTION_NAME         = "test_connection_name"
	TEST_SERVICE_CONNECTION_SERVICE_CODE = "test_connection_code"
)

func buildServiceConnectionJsonResponse(serviceConnection *ServiceConnection) []byte {
	j, _ := json.Marshal(serviceConnection)
	return j
}

func buildListServiceConnectionJsonResponse(serviceConnections []ServiceConnection) []byte {
	j, _ := json.Marshal(serviceConnections)
	return j
}

func TestGetServiceConnectionReturnServiceConnectionIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	serviceConnectionService := ServiceConnectionApi{
		configurationService: mockConfigurationService,
	}

	expectedServiceConnection := ServiceConnection{Id: TEST_SERVICE_CONNECTION_ID,
		Name:        TEST_SERVICE_CONNECTION_NAME,
		ServiceCode: TEST_SERVICE_CONNECTION_SERVICE_CODE}

	mockConfigurationService.EXPECT().Get(TEST_SERVICE_CONNECTION_ID, gomock.Any()).Return(buildServiceConnectionJsonResponse(&expectedServiceConnection), nil)

	//when
	serviceConnection, _ := serviceConnectionService.Get(TEST_SERVICE_CONNECTION_ID)

	//then
	if assert.NotNil(t, serviceConnection) {
		assert.Equal(t, expectedServiceConnection, *serviceConnection)
	}
}

func TestGetServiceConnectionReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	serviceConnectionService := ServiceConnectionApi{
		configurationService: mockConfigurationService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockConfigurationService.EXPECT().Get(TEST_SERVICE_CONNECTION_ID, gomock.Any()).Return(nil, mockError)

	//when
	serviceConnection, err := serviceConnectionService.Get(TEST_SERVICE_CONNECTION_ID)

	//then
	assert.Nil(t, serviceConnection)
	assert.Equal(t, mockError, err)

}

func TestListServiceConnectionReturnDiskOfferingsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	serviceConnectionService := ServiceConnectionApi{
		configurationService: mockConfigurationService,
	}

	expectedServiceConnections := []ServiceConnection{
		{
			Id:          "connection_1",
			Name:        "connection_name_1",
			ServiceCode: "connection_code_1",
		},
		{
			Id:          "connection_2",
			Name:        "connection_name_2",
			ServiceCode: "connection_code_2",
		},
	}

	mockConfigurationService.EXPECT().List(gomock.Any()).Return(buildListServiceConnectionJsonResponse(expectedServiceConnections), nil)

	//when
	serviceConnections, _ := serviceConnectionService.List()

	//then
	if assert.NotNil(t, serviceConnections) {
		assert.Equal(t, expectedServiceConnections, serviceConnections)
	}
}

func TestListServiceConnectionReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	serviceConnectionService := ServiceConnectionApi{
		configurationService: mockConfigurationService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockConfigurationService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	serviceConnections, err := serviceConnectionService.List()

	//then
	assert.Nil(t, serviceConnections)
	assert.Equal(t, mockError, err)

}
