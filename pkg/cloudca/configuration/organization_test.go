package configuration

import (
	"encoding/json"

	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/configuration_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	TEST_ORGANIZATION_ID           = "org_id"
	TEST_ORGANIZATION_NAME         = "test_org_name"
	TEST_ORGANIZATION_ENTRYPOINT   = "test_entrypoint"
	TEST_ORGANIZATION_USERS        = []User{{Id: "test_user1"}, {Id: "test_user2"}}
	TEST_ORGANIZATION_ENVIRONMENTS = []Environment{{Id: "test_env1"}, {Id: "test_env2"}}
	TEST_ORGANIZATION_ROLES        = []Role{{Id: "test_role"}}
)

func buildOrganizationJsonResponse(organization *Organization) []byte {
	j, _ := json.Marshal(organization)
	return j
}

func buildListOrganizationJsonResponse(organizations []Organization) []byte {
	j, _ := json.Marshal(organizations)
	return j
}

func TestGetOrganizationReturnOrganizationIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	organizationService := OrganizationApi{
		configurationService: mockConfigurationService,
	}

	expectedOrganization := Organization{Id: TEST_ORGANIZATION_ID,
		Name:         TEST_ORGANIZATION_NAME,
		EntryPoint:   TEST_ORGANIZATION_ENTRYPOINT,
		Users:        TEST_ORGANIZATION_USERS,
		Environments: TEST_ORGANIZATION_ENVIRONMENTS,
		Roles:        TEST_ORGANIZATION_ROLES}

	mockConfigurationService.EXPECT().Get(TEST_ORGANIZATION_ID, gomock.Any()).Return(buildOrganizationJsonResponse(&expectedOrganization), nil)

	//when
	organization, _ := organizationService.Get(TEST_ORGANIZATION_ID)

	//then
	if assert.NotNil(t, organization) {
		assert.Equal(t, expectedOrganization, *organization)
	}
}

func TestGetOrganizationReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	organizationService := OrganizationApi{
		configurationService: mockConfigurationService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockConfigurationService.EXPECT().Get(TEST_ORGANIZATION_ID, gomock.Any()).Return(nil, mockError)

	//when
	organization, err := organizationService.Get(TEST_ORGANIZATION_ID)

	//then
	assert.Nil(t, organization)
	assert.Equal(t, mockError, err)

}

func TestListOrganizationReturnDiskOfferingsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	organizationService := OrganizationApi{
		configurationService: mockConfigurationService,
	}

	expectedOrganizations := []Organization{
		{
			Id:           "org_id_1",
			Name:         "org_name_1",
			EntryPoint:   "org_entrypoint_1",
			Users:        []User{{Id: "user1"}},
			Environments: []Environment{},
			Roles:        []Role{{Id: "test_role_1"}},
		},
		{
			Id:           "org_id_2",
			Name:         "org_name_2",
			EntryPoint:   "org_entrypoint_2",
			Users:        []User{{Id: "user2"}},
			Environments: []Environment{{Id: "env1"}},
			Roles:        []Role{{Id: "test_role_2"}},
		},
	}

	mockConfigurationService.EXPECT().List(gomock.Any()).Return(buildListOrganizationJsonResponse(expectedOrganizations), nil)

	//when
	organizations, _ := organizationService.List()

	//then
	if assert.NotNil(t, organizations) {
		assert.Equal(t, expectedOrganizations, organizations)
	}
}

func TestListOrganizationReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)

	organizationService := OrganizationApi{
		configurationService: mockConfigurationService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockConfigurationService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	organizations, err := organizationService.List()

	//then
	assert.Nil(t, organizations)
	assert.Equal(t, mockError, err)

}
