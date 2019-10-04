package cloudca

import (
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_PUBLIC_IP_ID = "test_public_ip_id"
	TEST_IP_ADDRESS   = "172.31.3.208"
)

func buildTestPublicIpJsonResponse(publicIp *PublicIp) []byte {
	return []byte(`{"id":"` + publicIp.Id + `",` +
		` "ipAddress":"` + publicIp.IpAddress + `"}`)
}

func buildListTestPublicIpJsonResponse(publicIps []PublicIp) []byte {
	resp := `[`
	for i, t := range publicIps {
		resp += string(buildTestPublicIpJsonResponse(&t))
		if i != len(publicIps)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetPublicIpReturnPublicIpIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	expectedPublicIp := PublicIp{Id: TEST_PUBLIC_IP_ID,
		IpAddress: TEST_IP_ADDRESS,
	}

	mockEntityService.EXPECT().Get(TEST_PUBLIC_IP_ID, gomock.Any()).Return(buildTestPublicIpJsonResponse(&expectedPublicIp), nil)

	//when
	publicIp, _ := publicIpService.Get(TEST_PUBLIC_IP_ID)

	//then
	if assert.NotNil(t, publicIp) {
		assert.Equal(t, expectedPublicIp, *publicIp)
	}
}

func TestGetPublicIpReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_PUBLIC_IP_ID, gomock.Any()).Return(nil, mockError)

	//when
	publicIp, err := publicIpService.Get(TEST_PUBLIC_IP_ID)

	//then
	assert.Nil(t, publicIp)
	assert.Equal(t, mockError, err)

}

func TestListPublicIpReturnNetworksIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	expectedPublicIp1 := PublicIp{Id: "list_id_1",
		IpAddress: "list_ip_address_1"}

	expectedPublicIp2 := PublicIp{Id: "list_id_2",
		IpAddress: "list_ip_address_2"}

	expectedPublicIps := []PublicIp{expectedPublicIp1, expectedPublicIp2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestPublicIpJsonResponse(expectedPublicIps), nil)

	//when
	publicIps, _ := publicIpService.List()

	//then
	if assert.NotNil(t, publicIps) {
		assert.Equal(t, expectedPublicIps, publicIps)
	}
}

func TestListPublicIpReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	publicIps, err := publicIpService.List()

	//then
	assert.Nil(t, publicIps)
	assert.Equal(t, mockError, err)

}

func TestPublicIpReturnAcquiredPublicIpIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	publicIpToAcquire := PublicIp{VpcId: "vpcId"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{"id":"new_id", "ipAddress": "new_ip_address"}`), nil)

	//when
	acquiredPublicIp, _ := publicIpService.Acquire(publicIpToAcquire)

	//then
	if assert.NotNil(t, acquiredPublicIp) {
		assert.Equal(t, "new_id", acquiredPublicIp.Id)
		assert.Equal(t, "new_ip_address", acquiredPublicIp.IpAddress)
	}
}

func TestPublicIpReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_instance_error"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	publicIpToAcquire := PublicIp{VpcId: "vpcId"}

	//when
	acquiredPublicIp, err := publicIpService.Acquire(publicIpToAcquire)

	//then
	assert.Nil(t, acquiredPublicIp)
	assert.Equal(t, mockError, err)
}

func TestPublicIpReleaseReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Delete(TEST_PUBLIC_IP_ID, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := publicIpService.Release(TEST_PUBLIC_IP_ID)

	//then
	assert.True(t, success)
}

func TestPublicIpReleaseReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_purge_instance_error"}
	mockEntityService.EXPECT().Delete(TEST_PUBLIC_IP_ID, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := publicIpService.Release(TEST_PUBLIC_IP_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestPublicIpEnableStaticNatReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	publicIp := PublicIp{
		Id:          TEST_PUBLIC_IP_ID,
		PrivateIpId: "private_ip_id",
	}

	mockEntityService.EXPECT().Execute(TEST_PUBLIC_IP_ID, PUBLIC_IP_ENABLE_STATIC_NAT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := publicIpService.EnableStaticNat(publicIp)

	//then
	assert.True(t, success)
}

func TestPublicIpEnableStaticNatReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	publicIp := PublicIp{
		Id:          TEST_PUBLIC_IP_ID,
		PrivateIpId: "private_ip_id",
	}

	mockError := mocks.MockError{"some_purge_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_PUBLIC_IP_ID, PUBLIC_IP_ENABLE_STATIC_NAT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := publicIpService.EnableStaticNat(publicIp)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestPublicIpDisableStaticNatReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_PUBLIC_IP_ID, PUBLIC_IP_DISABLE_STATIC_NAT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := publicIpService.DisableStaticNat(TEST_PUBLIC_IP_ID)

	//then
	assert.True(t, success)
}

func TestPublicIpDisableStaticNatReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	publicIpService := PublicIpApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_purge_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_PUBLIC_IP_ID, PUBLIC_IP_DISABLE_STATIC_NAT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := publicIpService.DisableStaticNat(TEST_PUBLIC_IP_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}
