package configuration

import (
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks/configuration_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConfigurationType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// this is a weak test. It validates that the constant is the correct API endpoint,
	// because the buildUrl method is unaccessible
	assert.Equal(t, ENVIRONMENT_CONFIGURATION_TYPE, "environments")
}

func TestGetEnvironmentReturnEnvironmentIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)
	environmentService := EnvironmentApi{
		configurationService: mockConfigurationService,
	}

	response := `{"id":"13ca7410-9b4a-4fd7-ae2e-e5455b664faf",
                 "name":"patdev1",
                 "description":"pat dev 1",
                 "membership":"MANY_USERS",
                 "serviceConnection":{"id":"73983e63-e404-48aa-a89c-f41ca93af9cd","category":"IAAS","name":"patDev1","serviceCode":"dev1","type":"CloudCA"},
                 "organization":{"id":"4b5e5c55-7aea-48e4-9287-d63b36457c51","entryPoint":"pat","name":"Test"},
                 "users":[{"id":"062445f9-11d3-4e7b-9a84-908272a72250","userName":"pdube"}],
                 "roles":[{"id":"32a25a1e-0506-429f-a731-e8fcaaa01c4d","users":[],"isDefault":false,"name":"Read-only"},
                          {"id":"517b40e5-20a8-44f0-a5d0-06ed20ee4d43","users":[{"id":"062445f9-11d3-4e7b-9a84-908272a72250","userName":"pdube"}],"isDefault":false,"name":"Environment Admin"}
                         ],
                 "deleted":false,
                 "version":5}`

	expectedEnvironment := Environment{
		Id:          "13ca7410-9b4a-4fd7-ae2e-e5455b664faf",
		Name:        "patdev1",
		Description: "pat dev 1",
		ServiceConnection: ServiceConnection{
			Id:          "73983e63-e404-48aa-a89c-f41ca93af9cd",
			Name:        "patDev1",
			ServiceCode: "dev1",
		},
		Organization: Organization{
			Id:         "4b5e5c55-7aea-48e4-9287-d63b36457c51",
			Name:       "Test",
			EntryPoint: "pat",
		},
		Users: []User{
			{
				Id:       "062445f9-11d3-4e7b-9a84-908272a72250",
				Username: "pdube",
			},
		},
		Roles: []Role{
			{
				Id:    "32a25a1e-0506-429f-a731-e8fcaaa01c4d",
				Name:  "Read-only",
				Users: []User{},
			},
			{
				Id:   "517b40e5-20a8-44f0-a5d0-06ed20ee4d43",
				Name: "Environment Admin",
				Users: []User{
					{
						Id:       "062445f9-11d3-4e7b-9a84-908272a72250",
						Username: "pdube",
					},
				},
			},
		},
	}

	mockConfigurationService.EXPECT().Get("13ca7410-9b4a-4fd7-ae2e-e5455b664faf", gomock.Any()).Return([]byte(response), nil)

	//when
	environment, _ := environmentService.Get("13ca7410-9b4a-4fd7-ae2e-e5455b664faf")

	//then
	if assert.NotNil(t, environment) {
		assert.Equal(t, expectedEnvironment, *environment)
	}
}

func TestListEnvironmentReturnEnvironmentIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)
	environmentService := EnvironmentApi{
		configurationService: mockConfigurationService,
	}

	response := `[{"id":"13ca7410-9b4a-4fd7-ae2e-e5455b664faf",
                 "name":"patdev1",
                 "description":"pat dev 1",
                 "membership":"MANY_USERS",
                 "serviceConnection":{"id":"73983e63-e404-48aa-a89c-f41ca93af9cd","category":"IAAS","name":"patDev1","serviceCode":"dev1","type":"CloudCA"},
                 "organization":{"id":"4b5e5c55-7aea-48e4-9287-d63b36457c51","entryPoint":"pat","name":"Test"},
                 "users":[{"id":"062445f9-11d3-4e7b-9a84-908272a72250","userName":"pdube"}],
                 "roles":[{"id":"32a25a1e-0506-429f-a731-e8fcaaa01c4d","users":[],"isDefault":false,"name":"Read-only"},
                          {"id":"517b40e5-20a8-44f0-a5d0-06ed20ee4d43","users":[{"id":"062445f9-11d3-4e7b-9a84-908272a72250","userName":"pdube"}],"isDefault":false,"name":"Environment Admin"}
                         ],
                 "deleted":false,
                 "version":5}]`

	expectedEnvironments := []Environment{
		{
			Id:          "13ca7410-9b4a-4fd7-ae2e-e5455b664faf",
			Name:        "patdev1",
			Description: "pat dev 1",
			ServiceConnection: ServiceConnection{
				Id:          "73983e63-e404-48aa-a89c-f41ca93af9cd",
				Name:        "patDev1",
				ServiceCode: "dev1",
			},
			Organization: Organization{
				Id:         "4b5e5c55-7aea-48e4-9287-d63b36457c51",
				Name:       "Test",
				EntryPoint: "pat",
			},
			Users: []User{
				{
					Id:       "062445f9-11d3-4e7b-9a84-908272a72250",
					Username: "pdube",
				},
			},
			Roles: []Role{
				{
					Id:    "32a25a1e-0506-429f-a731-e8fcaaa01c4d",
					Name:  "Read-only",
					Users: []User{},
				},
				{
					Id:   "517b40e5-20a8-44f0-a5d0-06ed20ee4d43",
					Name: "Environment Admin",
					Users: []User{
						{
							Id:       "062445f9-11d3-4e7b-9a84-908272a72250",
							Username: "pdube",
						},
					},
				},
			},
		},
	}

	mockConfigurationService.EXPECT().List(gomock.Any()).Return([]byte(response), nil)

	//when
	environments, _ := environmentService.List()

	//then
	if assert.NotNil(t, environments) {
		assert.Equal(t, expectedEnvironments, environments)
	}
}
