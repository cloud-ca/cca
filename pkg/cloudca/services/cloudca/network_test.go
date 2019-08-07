package cloudca

import (
	"strconv"
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_NETWORK_ID                  = "test_network_id"
	TEST_NETWORK_NAME                = "test_network"
	TEST_NETWORK_ZONE_ID             = "test_network_zone_id"
	TEST_NETWORK_ZONE_NAME           = "test_network_zone_name"
	TEST_NETWORK_CIDR                = "test_network_cidr"
	TEST_NETWORK_TYPE                = "test_network_type"
	TEST_NETWORK_STATE               = "test_network_state"
	TEST_NETWORK_GATEWAY             = "test_network_gateway"
	TEST_NETWORK_NETWORK_OFFERING_ID = "test_network_network_offering_id"
	TEST_NETWORK_IS_SYSTEM           = false
	TEST_NETWORK_VPC_ID              = "test_network_vpc_id"
	TEST_NETWORK_DOMAIN              = "test_network_domain"
	TEST_NETWORK_DOMAIN_ID           = "test_network_domain_id"
	TEST_NETWORK_PROJECT             = "test_network_project"
	TEST_NETWORK_PROJECT_ID          = "test_network_project_id"
	TEST_NETWORK_ACL_ID_REF          = "test_network_acl_id"
)

func buildTestNetworkJsonResponse(network *Network) []byte {
	return []byte(`{"id":"` + network.Id + `",` +
		` "name":"` + network.Name + `",` +
		` "zoneid":"` + network.ZoneId + `",` +
		` "zonename":"` + network.ZoneName + `",` +
		` "cidr":"` + network.Cidr + `",` +
		` "type":"` + network.Type + `",` +
		` "state":"` + network.State + `",` +
		` "gateway":"` + network.Gateway + `",` +
		` "networkOfferingId":"` + network.NetworkOfferingId + `",` +
		` "issystem":` + strconv.FormatBool(network.IsSystem) + `,` +
		` "vpcId":"` + network.VpcId + `",` +
		` "domain":"` + network.Domain + `",` +
		` "domainid":"` + network.DomainId + `",` +
		` "project":"` + network.Project + `",` +
		` "projectid":"` + network.ProjectId + `",` +
		` "networkACLId":"` + network.NetworkAclId + `"}`)
}

func buildListTestNetworkJsonResponse(networks []Network) []byte {
	resp := `[`
	for i, t := range networks {
		resp += string(buildTestNetworkJsonResponse(&t))
		if i != len(networks)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetNetworkReturnNetworkIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkService := NetworkApi{
		entityService: mockEntityService,
	}

	expectedNetwork := Network{Id: TEST_NETWORK_ID,
		Name:              TEST_NETWORK_NAME,
		ZoneId:            TEST_NETWORK_ZONE_ID,
		ZoneName:          TEST_NETWORK_ZONE_NAME,
		Cidr:              TEST_NETWORK_CIDR,
		Type:              TEST_NETWORK_TYPE,
		Gateway:           TEST_NETWORK_GATEWAY,
		NetworkOfferingId: TEST_NETWORK_NETWORK_OFFERING_ID,
		IsSystem:          TEST_NETWORK_IS_SYSTEM,
		VpcId:             TEST_NETWORK_VPC_ID,
		Domain:            TEST_NETWORK_DOMAIN,
		DomainId:          TEST_NETWORK_DOMAIN_ID,
		Project:           TEST_NETWORK_PROJECT,
		ProjectId:         TEST_NETWORK_PROJECT_ID,
		NetworkAclId:      TEST_NETWORK_ACL_ID_REF}

	mockEntityService.EXPECT().Get(TEST_NETWORK_ID, gomock.Any()).Return(buildTestNetworkJsonResponse(&expectedNetwork), nil)

	//when
	network, _ := networkService.Get(TEST_NETWORK_ID)

	//then
	if assert.NotNil(t, network) {
		assert.Equal(t, expectedNetwork, *network)
	}
}

func TestGetNetworkReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkService := NetworkApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_NETWORK_ID, gomock.Any()).Return(nil, mockError)

	//when
	network, err := networkService.Get(TEST_NETWORK_ID)

	//then
	assert.Nil(t, network)
	assert.Equal(t, mockError, err)

}

func TestListNetworkReturnNetworksIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkService := NetworkApi{
		entityService: mockEntityService,
	}

	expectedNetwork1 := Network{Id: "list_id_1",
		Name:              "list_name_1",
		ZoneId:            "list_zone_id_1",
		ZoneName:          "list_zone_name_1",
		Cidr:              "list_cidr_1",
		Type:              "list_type_1",
		Gateway:           "list_gateway_1",
		NetworkOfferingId: "list_network_offering_id_1",
		IsSystem:          true,
		VpcId:             "list_vpc_id_1",
		Domain:            "list_domain_1",
		DomainId:          "list_domain_id_1",
		Project:           "list_project_1",
		ProjectId:         "list_project_id_1",
		NetworkAclId:      "list_acl_id_1"}

	expectedNetwork2 := Network{Id: "list_id_2",
		Name:              "list_name_2",
		ZoneId:            "list_zone_id_2",
		ZoneName:          "list_zone_name_2",
		Cidr:              "list_cidr_2",
		Type:              "list_type_2",
		Gateway:           "list_gateway_2",
		NetworkOfferingId: "list_network_offering_id_2",
		IsSystem:          false,
		VpcId:             "list_vpc_id_2",
		Domain:            "list_domain_2",
		DomainId:          "list_domain_id_2",
		Project:           "list_project_2",
		ProjectId:         "list_project_id_2",
		NetworkAclId:      "list_acl_id_2"}

	expectedNetworks := []Network{expectedNetwork1, expectedNetwork2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestNetworkJsonResponse(expectedNetworks), nil)

	//when
	networks, _ := networkService.List()

	//then
	if assert.NotNil(t, networks) {
		assert.Equal(t, expectedNetworks, networks)
	}
}

func TestListNetworkReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkService := NetworkApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	networks, err := networkService.List()

	//then
	assert.Nil(t, networks)
	assert.Equal(t, mockError, err)

}
