package cloudca

import (
	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TEST_VPC_ID              = "test_vpc_id"
	TEST_VPC_NAME            = "test_vpc"
	TEST_VPC_DESCRIPTION     = "test_vpc_description"
	TEST_VPC_STATE           = "test_vpc_state"
	TEST_VPC_CIDR            = "test_vpc_cidr"
	TEST_VPC_ZONE_ID         = "test_vpc_zone_id"
	TEST_VPC_ZONE_NAME       = "test_vpc_zone_name"
	TEST_VPC_NETWORK_DOMAIN  = "test_vpc_network_domain"
	TEST_VPC_SOURCE_NAT_IP   = "test_vpc_source_nat_ip"
	TEST_VPC_VPN_STATUS      = "test_vpc_vpn_status"
	TEST_VPC_TYPE            = "test_vpc_type"
	TEST_VPC_VPC_OFFERING_ID = "test_vpc_offering_id"
)

func buildTestVpcJsonResponse(vpc *Vpc) []byte {
	return []byte(`{"id": "` + vpc.Id + `", ` +
		`"name":"` + vpc.Name + `", ` +
		`"description":"` + vpc.Description + `", ` +
		`"state":"` + vpc.State + `", ` +
		`"cidr":"` + vpc.Cidr + `", ` +
		`"zoneId":"` + vpc.ZoneId + `", ` +
		`"zoneName":"` + vpc.ZoneName + `", ` +
		`"networkDomain":"` + vpc.NetworkDomain + `", ` +
		`"sourceNatIp":"` + vpc.SourceNatIp + `", ` +
		`"vpnStatus":"` + vpc.VpnStatus + `", ` +
		`"type":"` + vpc.Type + `", ` +
		`"vpcOfferingId":"` + vpc.VpcOfferingId + `"}`)
}

func buildListTestVpcJsonResponse(vpcs []Vpc) []byte {
	resp := `[`
	for i, vpc := range vpcs {
		resp += string(buildTestVpcJsonResponse(&vpc))
		if i != len(vpcs)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetVpcReturnVpcIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	expectedVpc := Vpc{Id: TEST_VPC_ID,
		Name:          TEST_VPC_NAME,
		Description:   TEST_VPC_DESCRIPTION,
		State:         TEST_VPC_STATE,
		Cidr:          TEST_VPC_CIDR,
		ZoneId:        TEST_VPC_ZONE_ID,
		ZoneName:      TEST_VPC_ZONE_NAME,
		NetworkDomain: TEST_VPC_NETWORK_DOMAIN,
		SourceNatIp:   TEST_VPC_SOURCE_NAT_IP,
		VpnStatus:     TEST_VPC_VPN_STATUS,
		Type:          TEST_VPC_TYPE,
		VpcOfferingId: TEST_VPC_VPC_OFFERING_ID,
	}

	mockEntityService.EXPECT().Get(TEST_VPC_ID, gomock.Any()).Return(buildTestVpcJsonResponse(&expectedVpc), nil)

	//when
	vpc, _ := vpcService.Get(TEST_VPC_ID)

	//then
	if assert.NotNil(t, vpc) {
		assert.Equal(t, expectedVpc, *vpc)
	}
}

func TestGetVpcReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_VPC_ID, gomock.Any()).Return(nil, mockError)

	//when
	vpc, err := vpcService.Get(TEST_VPC_ID)

	//then
	assert.Nil(t, vpc)
	assert.Equal(t, mockError, err)

}

func TestListVpcReturnVpcsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	expectedVpc1 := Vpc{Id: TEST_VPC_ID + "1",
		Name:          TEST_VPC_NAME + "1",
		Description:   TEST_VPC_DESCRIPTION + "1",
		State:         TEST_VPC_STATE + "1",
		Cidr:          TEST_VPC_CIDR + "1",
		ZoneId:        TEST_VPC_ZONE_ID + "1",
		ZoneName:      TEST_VPC_ZONE_NAME + "1",
		NetworkDomain: TEST_VPC_NETWORK_DOMAIN + "1",
		SourceNatIp:   TEST_VPC_SOURCE_NAT_IP + "1",
		VpnStatus:     TEST_VPC_VPN_STATUS + "1",
		Type:          TEST_VPC_TYPE + "1",
		VpcOfferingId: TEST_VPC_VPC_OFFERING_ID + "1",
	}

	expectedVpcs := []Vpc{expectedVpc1}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestVpcJsonResponse(expectedVpcs), nil)

	//when
	vpcs, _ := vpcService.List()

	//then
	if assert.NotNil(t, vpcs) {
		assert.Equal(t, expectedVpcs, vpcs)
	}
}

func TestListVpcReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	vpcs, err := vpcService.List()

	//then
	assert.Nil(t, vpcs)
	assert.Equal(t, mockError, err)

}

func TestCreateVpcReturnCreatedInstanceIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	vpcToCreate := Vpc{Name: "new_name",
		Description:   "new_description",
		VpcOfferingId: "vpc_offering_id",
	}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{"id":"new_id"}`), nil)

	//when
	createdVpc, _ := vpcService.Create(vpcToCreate)

	//then
	if assert.NotNil(t, createdVpc) {
		assert.Equal(t, "new_id", createdVpc.Id)
	}
}

func TestCreateVpcReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_vpc_error"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	vpcToCreate := Vpc{Name: "new_name",
		Description:   "new_description",
		VpcOfferingId: "vpcOfferingId"}

	//when
	createdVpc, err := vpcService.Create(vpcToCreate)

	//then
	assert.Nil(t, createdVpc)
	assert.Equal(t, mockError, err)

}

func TestRestartRouterReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_VPC_ID, VPC_RESTART_ROUTER_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := vpcService.RestartRouter(TEST_VPC_ID)

	//then
	assert.True(t, success)
}

func TestRestartRouterReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcService := VpcApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_purge_vpc_error"}
	mockEntityService.EXPECT().Execute(TEST_VPC_ID, VPC_RESTART_ROUTER_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := vpcService.RestartRouter(TEST_VPC_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}
