package cloudca

import (
	"fmt"
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const PFR_TEMPLATE = `{
	"id": "%s",
	"instanceId": "1",
	"instanceName": "instance0",
	"networkId": "2",
	"privateIp": "127.0.0.1",
	"privateIpId": "3",
	"privatePortStart": "8080",
	"privatePortEnd": "8080",
	"ipAddress": "192.168.0.1",
	"ipAddressId": "4",
	"publicPortStart": "80",
	"publicPortEnd": "80",
	"protocol": "TCP",
	"state": "Active",
	"vpcId": "5"
}`

func setupMock(t *testing.T) *services_mocks.MockEntityService {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return services_mocks.NewMockEntityService(ctrl)
}

func createPfrWithId(id string) *PortForwardingRule {
	return &PortForwardingRule{
		Id:               id,
		InstanceId:       "1",
		InstanceName:     "instance0",
		NetworkId:        "2",
		PrivateIp:        "127.0.0.1",
		PrivateIpId:      "3",
		PrivatePortStart: "8080",
		PrivatePortEnd:   "8080",
		PublicIp:         "192.168.0.1",
		PublicIpId:       "4",
		PublicPortStart:  "80",
		PublicPortEnd:    "80",
		Protocol:         "TCP",
		State:            "Active",
		VpcId:            "5",
	}
}

func TestGetById(t *testing.T) {
	// given
	mockEntityService := setupMock(t)
	pfrService := PortForwardingRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "pfr_0"
	expectedPfr := *createPfrWithId(expectedId)

	response := fmt.Sprintf(PFR_TEMPLATE, expectedId)
	mockEntityService.EXPECT().Get(expectedId, gomock.Any()).Return([]byte(response), nil)

	// when
	pfr, _ := pfrService.Get(expectedId)

	// then
	assert.Equal(t, expectedPfr, *pfr)
}

func TestListWithOptions(t *testing.T) {
	// given
	mockEntityService := setupMock(t)
	pfrService := PortForwardingRuleApi{
		entityService: mockEntityService,
	}

	id1, id2 := "1234", "4321"
	pfr1, pfr2 := fmt.Sprintf(PFR_TEMPLATE, id1), fmt.Sprintf(PFR_TEMPLATE, id2)
	response := fmt.Sprintf("[ %s, %s ]", pfr1, pfr2)
	mockEntityService.EXPECT().List(gomock.Any()).Return([]byte(response), nil)

	// when
	pfrs, _ := pfrService.ListWithOptions(map[string]string{})

	// then
	assert.Equal(t, id1, pfrs[0].Id)
	assert.Equal(t, id2, pfrs[1].Id)
}

func TestCreate(t *testing.T) {
	// given
	mockEntityService := setupMock(t)
	pfrService := PortForwardingRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "adsf"
	response := fmt.Sprintf(PFR_TEMPLATE, expectedId)
	expectedPfr := *createPfrWithId(expectedId)

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(response), nil)

	// when
	pfr, _ := pfrService.Create(expectedPfr)

	// then
	assert.Equal(t, expectedPfr, *pfr)
}

func TestDeleteReturnsSuccess_ifNoErrorsOccur(t *testing.T) {
	// given
	mockEntityService := setupMock(t)
	pfrService := PortForwardingRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "id0"
	mockEntityService.EXPECT().Delete(expectedId, gomock.Any(), gomock.Any()).Return([]byte{}, nil)

	// when
	success, _ := pfrService.Delete(expectedId)

	// then
	assert.True(t, success)
}

func TestDeleteReturnsFailure_ifErrorOccurred(t *testing.T) {
	// given
	mockEntityService := setupMock(t)
	pfrService := PortForwardingRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "id0"
	mockEntityService.EXPECT().Delete(expectedId, gomock.Any(), gomock.Any()).Return(nil, mocks.MockError{"asdf"})

	// when
	success, _ := pfrService.Delete(expectedId)

	// then
	assert.False(t, success)
}
