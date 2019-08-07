package cloudca

import (
	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TEST_SSH_KEY_NAME        = "test_ssh_key"
	TEST_SSH_KEY_FINGERPRINT = "test_fingerprint"
)

func buildSSHKeyJsonResponse(sshKey *SSHKey) []byte {
	return []byte(`{"name": "` + sshKey.Name +
		`","fingerprint":"` + sshKey.Fingerprint + `"}`)
}

func buildListSSHKeyJsonResponse(sshKeys []SSHKey) []byte {
	resp := `[`
	for i, s := range sshKeys {
		resp += string(buildSSHKeyJsonResponse(&s))
		if i != len(sshKeys)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetSSHKeyReturnSSHKeyIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	expectedSSHKey := SSHKey{Name: TEST_SSH_KEY_NAME,
		Fingerprint: TEST_SSH_KEY_FINGERPRINT}

	mockEntityService.EXPECT().Get(TEST_SSH_KEY_NAME, gomock.Any()).Return(buildSSHKeyJsonResponse(&expectedSSHKey), nil)

	//when
	sshKey, _ := sshKeyService.Get(TEST_SSH_KEY_NAME)

	//then
	if assert.NotNil(t, sshKey) {
		assert.Equal(t, expectedSSHKey, *sshKey)
	}
}

func TestGetSSHKeyReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_SSH_KEY_NAME, gomock.Any()).Return(nil, mockError)

	//when
	sshKey, err := sshKeyService.Get(TEST_SSH_KEY_NAME)

	//then
	assert.Nil(t, sshKey)
	assert.Equal(t, mockError, err)

}

func TestListSSHKeyReturnSSHKeysIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	expectedSSHKeys := []SSHKey{
		{
			Name:        "list_name_1",
			Fingerprint: "list_fingerprint_1",
		},
		{
			Name:        "list_name_2",
			Fingerprint: "list_fingerprint_2",
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListSSHKeyJsonResponse(expectedSSHKeys), nil)

	//when
	sshKeys, _ := sshKeyService.List()

	//then
	if assert.NotNil(t, sshKeys) {
		assert.Equal(t, expectedSSHKeys, sshKeys)
	}
}

func TestListSSHKeyReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	sshKeys, err := sshKeyService.List()

	//then
	assert.Nil(t, sshKeys)
	assert.Equal(t, mockError, err)

}
