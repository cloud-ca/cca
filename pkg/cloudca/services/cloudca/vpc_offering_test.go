package cloudca

import (
	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TEST_VPC_OFFERING_ID    = "test_vpc_offering_id"
	TEST_VPC_OFFERING_NAME  = "test_vpc_offering"
	TEST_VPC_OFFERING_STATE = "test_vpc_offering_state"
)

func buildTestVpcOfferingJsonResponse(vpcOffering *VpcOffering) []byte {
	return []byte(`{"id": "` + vpcOffering.Id + `", ` +
		`"name":"` + vpcOffering.Name + `", ` +
		`"state":"` + vpcOffering.State + `"}`)
}

func buildListTestVpcOfferingJsonResponse(vpcOfferings []VpcOffering) []byte {
	resp := `[`
	for i, vpcOffering := range vpcOfferings {
		resp += string(buildTestVpcOfferingJsonResponse(&vpcOffering))
		if i != len(vpcOfferings)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetVpcOfferingReturnVpcOfferingIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcOfferingService := VpcOfferingApi{
		entityService: mockEntityService,
	}

	expectedVpcOffering := VpcOffering{Id: TEST_VPC_OFFERING_ID,
		Name:  TEST_VPC_OFFERING_NAME,
		State: TEST_VPC_OFFERING_STATE,
	}

	mockEntityService.EXPECT().Get(TEST_VPC_OFFERING_ID, gomock.Any()).Return(buildTestVpcOfferingJsonResponse(&expectedVpcOffering), nil)

	//when
	vpcOffering, _ := vpcOfferingService.Get(TEST_VPC_OFFERING_ID)

	//then
	if assert.NotNil(t, vpcOffering) {
		assert.Equal(t, expectedVpcOffering, *vpcOffering)
	}
}

func TestGetVpcOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcOfferingService := VpcOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_VPC_OFFERING_ID, gomock.Any()).Return(nil, mockError)

	//when
	vpcOffering, err := vpcOfferingService.Get(TEST_VPC_OFFERING_ID)

	//then
	assert.Nil(t, vpcOffering)
	assert.Equal(t, mockError, err)

}

func TestListVpcOfferingReturnVpcsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcOfferingService := VpcOfferingApi{
		entityService: mockEntityService,
	}

	expectedVpcOffering1 := VpcOffering{Id: TEST_VPC_OFFERING_ID + "1",
		Name:  TEST_VPC_OFFERING_NAME + "1",
		State: TEST_VPC_OFFERING_STATE + "1",
	}

	expectedVpcOfferings := []VpcOffering{expectedVpcOffering1}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestVpcOfferingJsonResponse(expectedVpcOfferings), nil)

	//when
	vpcOfferings, _ := vpcOfferingService.List()

	//then
	if assert.NotNil(t, vpcOfferings) {
		assert.Equal(t, expectedVpcOfferings, vpcOfferings)
	}
}

func TestListVpcOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	vpcOfferingService := VpcOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	vpcOfferings, err := vpcOfferingService.List()

	//then
	assert.Nil(t, vpcOfferings)
	assert.Equal(t, mockError, err)

}
