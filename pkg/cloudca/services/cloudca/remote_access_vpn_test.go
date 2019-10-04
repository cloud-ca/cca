package cloudca

import (
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_VPN_CERTIFICATE          = "test_vpn_certificate"
	TEST_VPN_ID                   = "test_vpn_id"
	TEST_VPN_PRESHARED_KEY        = "test_vpn_preshared_key"
	TEST_VPN_PUBLIC_IP_ADDRESS    = "test_vpn_public_ip_address"
	TEST_VPN_PUBLIC_IP_ADDRESS_ID = "test_vpn_public_ip_address_id"
	TEST_VPN_STATE                = "test_vpn_state"
	TEST_VPN_TYPE                 = "test_vpn_type"
)

func buildTestRemoteAccessVpnJsonResponse(remoteAccessVpn *RemoteAccessVpn) []byte {
	return []byte(`{"certificate":"` + remoteAccessVpn.Certificate + `",` +
		` "id":"` + remoteAccessVpn.Id + `",` +
		` "presharedKey":"` + remoteAccessVpn.PresharedKey + `",` +
		` "publicIpAddress":"` + remoteAccessVpn.PublicIpAddress + `",` +
		` "publicIpAddressId":"` + remoteAccessVpn.PublicIpAddressId + `",` +
		` "state":"` + remoteAccessVpn.State + `",` +
		` "type":"` + remoteAccessVpn.Type + `"}`)
}

func buildListTestRemoteAccessVpnJsonResponse(remoteAccessVpns []RemoteAccessVpn) []byte {
	resp := `[`
	for i, remoteAccessVpn := range remoteAccessVpns {
		resp += string(buildTestRemoteAccessVpnJsonResponse(&remoteAccessVpn))
		if i != len(remoteAccessVpns)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetRemoteAccessVpnReturnVpnIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	expectedRemoteAccessVpn := RemoteAccessVpn{
		Certificate:       TEST_VPN_CERTIFICATE,
		Id:                TEST_VPN_ID,
		PresharedKey:      TEST_VPN_PRESHARED_KEY,
		PublicIpAddress:   TEST_VPN_PUBLIC_IP_ADDRESS,
		PublicIpAddressId: TEST_VPN_PUBLIC_IP_ADDRESS_ID,
		State:             TEST_VPN_STATE,
		Type:              TEST_VPN_TYPE,
	}

	mockEntityService.EXPECT().Get(TEST_VPN_ID, gomock.Any()).Return(buildTestRemoteAccessVpnJsonResponse(&expectedRemoteAccessVpn), nil)

	//when
	remoteAccessVpn, _ := remoteAccessVpnService.Get(TEST_VPN_ID)

	//then
	if assert.NotNil(t, remoteAccessVpn) {
		assert.Equal(t, expectedRemoteAccessVpn, *remoteAccessVpn)
	}
}

func TestGetRemoteAccessVpnReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_VPN_ID, gomock.Any()).Return(nil, mockError)

	//when
	remoteAccessVpn, err := remoteAccessVpnService.Get(TEST_VPN_ID)

	//then
	assert.Nil(t, remoteAccessVpn)
	assert.Equal(t, mockError, err)

}

func TestListRemoteAccessVpnReturnVpnsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	expectedRemoteAccessVpn := RemoteAccessVpn{
		Certificate:       TEST_VPN_CERTIFICATE,
		Id:                TEST_VPN_ID,
		PresharedKey:      TEST_VPN_PRESHARED_KEY,
		PublicIpAddress:   TEST_VPN_PUBLIC_IP_ADDRESS,
		PublicIpAddressId: TEST_VPN_PUBLIC_IP_ADDRESS_ID,
		State:             TEST_VPN_STATE,
		Type:              TEST_VPN_TYPE,
	}

	expectedRemoteAccessVpns := []RemoteAccessVpn{expectedRemoteAccessVpn}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestRemoteAccessVpnJsonResponse(expectedRemoteAccessVpns), nil)

	//when
	remoteAccessVpns, _ := remoteAccessVpnService.List()

	//then
	if assert.NotNil(t, remoteAccessVpns) {
		assert.Equal(t, expectedRemoteAccessVpns, remoteAccessVpns)
	}
}

func TestListRemoteAccessVpnReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	remoteAccessVpns, err := remoteAccessVpnService.List()

	//then
	assert.Nil(t, remoteAccessVpns)
	assert.Equal(t, mockError, err)

}

func TestEnableRemoteAccessVpnReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_VPN_ID, REMOTE_ACCESS_VPN_ENABLE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := remoteAccessVpnService.Enable(TEST_VPN_ID)

	//then
	assert.True(t, success)
}

func TestEnableRemoteAccessVpnReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_vpn_enable_error"}
	mockEntityService.EXPECT().Execute(TEST_VPN_ID, REMOTE_ACCESS_VPN_ENABLE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := remoteAccessVpnService.Enable(TEST_VPN_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestDisableRemoteAccessVpnReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_VPN_ID, REMOTE_ACCESS_VPN_DISABLE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := remoteAccessVpnService.Disable(TEST_VPN_ID)

	//then
	assert.True(t, success)
}

func TestDisableRemoteAccessVpnReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnService := RemoteAccessVpnApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_vpn_disable_error"}
	mockEntityService.EXPECT().Execute(TEST_VPN_ID, REMOTE_ACCESS_VPN_DISABLE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := remoteAccessVpnService.Disable(TEST_VPN_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}
