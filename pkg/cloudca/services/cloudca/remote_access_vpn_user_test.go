package cloudca

import (
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_VPN_USER_ID       = "test_vpn_user_id"
	TEST_VPN_USER_USERNAME = "test_vpn_user_username"
	TEST_VPN_USER_PASSWORD = "test_vpn_user_password"
)

func buildTestRemoteAccessVpnUserJsonResponse(remoteAccessVpnUser *RemoteAccessVpnUser) []byte {
	return []byte(`{` +
		` "id":"` + remoteAccessVpnUser.Id + `",` +
		` "username":"` + remoteAccessVpnUser.Username + `"}`)
}

func buildListTestRemoteAccessVpnUserJsonResponse(remoteAccessVpnUsers []RemoteAccessVpnUser) []byte {
	resp := `[`
	for i, remoteAccessVpnUser := range remoteAccessVpnUsers {
		resp += string(buildTestRemoteAccessVpnUserJsonResponse(&remoteAccessVpnUser))
		if i != len(remoteAccessVpnUsers)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetRemoteAccessVpnUserReturnVpnUserIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	expectedRemoteAccessVpnUser := RemoteAccessVpnUser{
		Id:       TEST_VPN_USER_ID,
		Username: TEST_VPN_USER_USERNAME,
	}

	mockEntityService.EXPECT().Get(TEST_VPN_USER_ID, gomock.Any()).Return(buildTestRemoteAccessVpnUserJsonResponse(&expectedRemoteAccessVpnUser), nil)

	//when
	remoteAccessVpnUser, _ := remoteAccessVpnUserService.Get(TEST_VPN_USER_ID)

	//then
	if assert.NotNil(t, remoteAccessVpnUser) {
		assert.Equal(t, expectedRemoteAccessVpnUser, *remoteAccessVpnUser)
	}
}

func TestGetRemoteAccessVpnUserReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_VPN_USER_ID, gomock.Any()).Return(nil, mockError)

	//when
	remoteAccessVpnUser, err := remoteAccessVpnUserService.Get(TEST_VPN_USER_ID)

	//then
	assert.Nil(t, remoteAccessVpnUser)
	assert.Equal(t, mockError, err)

}

func TestListRemoteAccessVpnUsersReturnVpnUsersIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	expectedRemoteAccessVpnUser := RemoteAccessVpnUser{
		Id:       TEST_VPN_USER_ID,
		Username: TEST_VPN_USER_USERNAME,
	}

	expectedRemoteAccessVpnUsers := []RemoteAccessVpnUser{expectedRemoteAccessVpnUser}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestRemoteAccessVpnUserJsonResponse(expectedRemoteAccessVpnUsers), nil)

	//when
	remoteAccessVpnUsers, _ := remoteAccessVpnUserService.List()

	//then
	if assert.NotNil(t, remoteAccessVpnUsers) {
		assert.Equal(t, expectedRemoteAccessVpnUsers, remoteAccessVpnUsers)
	}
}

func TestListRemoteAccessVpnUsersReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	remoteAccessVpnUsers, err := remoteAccessVpnUserService.List()

	//then
	assert.Nil(t, remoteAccessVpnUsers)
	assert.Equal(t, mockError, err)
}

func TestCreateRemoteAccessVpnUserReturnVpnUserIfSuccess(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	createRemoteAccessVpnUser := RemoteAccessVpnUser{
		Username: TEST_VPN_USER_USERNAME,
		Password: TEST_VPN_USER_PASSWORD,
	}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	// when
	success, _ := remoteAccessVpnUserService.Create(createRemoteAccessVpnUser)

	// then
	assert.True(t, success)
}

func TestCreateRemoteAccessVpnUserReturnNilWithErrorIfError(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_error"}
	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	success, err := remoteAccessVpnUserService.Create(RemoteAccessVpnUser{})

	// then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestDeleteRemoteAccessVpnUserReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	deleteRemoteAccessVpnUser := RemoteAccessVpnUser{
		Username: TEST_VPN_USER_USERNAME,
		Password: TEST_VPN_USER_PASSWORD,
	}

	mockEntityService.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := remoteAccessVpnUserService.Delete(deleteRemoteAccessVpnUser)

	//then
	assert.True(t, success)
}

func TestDeleteRemoteAccessVpnUserReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	remoteAccessVpnUserService := RemoteAccessVpnUserApi{
		entityService: mockEntityService,
	}

	deleteRemoteAccessVpnUser := RemoteAccessVpnUser{
		Username: TEST_VPN_USER_USERNAME,
		Password: TEST_VPN_USER_PASSWORD,
	}

	mockError := mocks.MockError{"some_delete_error"}
	mockEntityService.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)
	//when
	success, err := remoteAccessVpnUserService.Delete(deleteRemoteAccessVpnUser)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}
