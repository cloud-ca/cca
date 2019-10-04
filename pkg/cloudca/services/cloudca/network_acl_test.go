package cloudca

import (
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_NETWORK_ACL_ID          = "test_network_acl_id"
	TEST_NETWORK_ACL_NAME        = "test_network_acl"
	TEST_NETWORK_ACL_DESCRIPTION = "test_network_acl_description"
	TEST_NETWORK_ACL_VPC_ID      = "test_network_acl_vpc_id"
)

func buildTestNetworkAclJsonResponse(networkAcl *NetworkAcl) []byte {
	return []byte(`{"id":"` + networkAcl.Id + `",` +
		` "name":"` + networkAcl.Name + `",` +
		` "description":"` + networkAcl.Description + `",` +
		` "vpcId":"` + networkAcl.VpcId + `"}`)
}

func buildListTestNetworkAclsJsonResponse(networkAcls []NetworkAcl) []byte {
	resp := `[`
	for i, t := range networkAcls {
		resp += string(buildTestNetworkAclJsonResponse(&t))
		if i != len(networkAcls)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetNetworkAclReturnNetworkAclIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	expectedNetworkAcl := NetworkAcl{Id: TEST_NETWORK_ACL_ID,
		Name:        TEST_NETWORK_ACL_NAME,
		Description: TEST_NETWORK_ACL_DESCRIPTION,
		VpcId:       TEST_NETWORK_ACL_VPC_ID}
	mockEntityService.EXPECT().Get(TEST_NETWORK_ACL_ID, gomock.Any()).Return(buildTestNetworkAclJsonResponse(&expectedNetworkAcl), nil)

	//when
	networkAcl, _ := networkAclService.Get(TEST_NETWORK_ACL_ID)

	//then
	if assert.NotNil(t, networkAcl) {
		assert.Equal(t, expectedNetworkAcl, *networkAcl)
	}
}

func TestGetNetworkAclReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_NETWORK_ACL_ID, gomock.Any()).Return(nil, mockError)

	//when
	networkAcl, err := networkAclService.Get(TEST_NETWORK_ACL_ID)

	//then
	assert.Nil(t, networkAcl)
	assert.Equal(t, mockError, err)

}

func TestListNetworkAclReturnNetworkAclsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	expectedNetworkAcl1 := NetworkAcl{Id: "list_id_1",
		Name:        "list_name_1",
		Description: "list_description_1",
		VpcId:       "list_vpc_id_1"}

	expectedNetworkAcl2 := NetworkAcl{Id: "list_id_2",
		Name:        "list_name_2",
		Description: "list_description_2",
		VpcId:       "list_vpc_id_2"}

	expectedNetworkAcls := []NetworkAcl{expectedNetworkAcl1, expectedNetworkAcl2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestNetworkAclsJsonResponse(expectedNetworkAcls), nil)

	//when
	networkAcls, _ := networkAclService.List()

	//then
	if assert.NotNil(t, networkAcls) {
		assert.Equal(t, expectedNetworkAcls, networkAcls)
	}
}

func TestListNetworkAclReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	networkAcls, err := networkAclService.List()

	//then
	assert.Nil(t, networkAcls)
	assert.Equal(t, mockError, err)

}

func TestCreateNetworkAclReturnCreatedNetworkAclIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	networkAclToCreate := NetworkAcl{Name: "new_name",
		Description: "new_description",
		VpcId:       "new_vpc",
	}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{"id":"new_id"}`), nil)

	//when
	createdNetworkAcl, _ := networkAclService.Create(networkAclToCreate)

	//then
	if assert.NotNil(t, createdNetworkAcl) {
		assert.Equal(t, "new_id", createdNetworkAcl.Id)
	}
}

func TestCreateNetworkAclReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_vpc_error"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	networkAclToCreate := NetworkAcl{Name: "new_name",
		Description: "new_description",
		VpcId:       "vpcId"}

	//when
	createdNetworkAcl, err := networkAclService.Create(networkAclToCreate)

	//then
	assert.Nil(t, createdNetworkAcl)
	assert.Equal(t, mockError, err)

}

func TestDeleteReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := networkAclService.Delete(TEST_NETWORK_ACL_ID)

	//then
	assert.True(t, success)
}

func TestDeleteReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	networkAclService := NetworkAclApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_delete_network_acl_id_error"}
	mockEntityService.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := networkAclService.Delete(TEST_VPC_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}
