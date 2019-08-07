package cloudca

import (
	"strconv"
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/api"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_INSTANCE_ID                    = "test_instance_id"
	TEST_INSTANCE_NAME                  = "test_instance"
	TEST_INSTANCE_STATE                 = "test_instance_state"
	TEST_INSTANCE_TEMPLATE_ID           = "test_instance_template_id"
	TEST_INSTANCE_TEMPLATE_NAME         = "test_instance_template_name"
	TEST_INSTANCE_IS_PASSWORD_ENABLED   = true
	TEST_INSTANCE_IS_SSH_KEY_ENABLED    = false
	TEST_INSTANCE_USERNAME              = "test_instance_username"
	TEST_INSTANCE_COMPUTE_OFFERING_ID   = "test_instance_compute_offering_id"
	TEST_INSTANCE_COMPUTE_OFFERING_NAME = "test_instance_compute_offering_name"
	TEST_INSTANCE_CPU_NUMBER            = 2
	TEST_INSTANCE_MEMORY_IN_MB          = 8000
	TEST_INSTANCE_ZONE_ID               = "test_instance_zone_id"
	TEST_INSTANCE_ZONE_NAME             = "test_instance_zone_name"
	TEST_INSTANCE_PROJECT_ID            = "test_instance_project_id"
	TEST_INSTANCE_NETWORK_ID            = "test_instance_network_id"
	TEST_INSTANCE_NETWORK_NAME          = "test_instance_network_name"
	TEST_INSTANCE_VPC_ID                = "test_instance_vpc_id"
	TEST_INSTANCE_VPC_NAME              = "test_instance_vpc_name"
	TEST_INSTANCE_MAC_ADDRESS           = "test_instance_mac_address"
	TEST_INSTANCE_IP_ADDRESS            = "test_instance_ip_address"
	TEST_INSTANCE_VOLUME_ID_TO_ATTACH   = "test_volume_id_to_attach"
	TEST_INSTANCE_USER_DATA             = "test_instance_user_data"
	TEST_INSTANCE_PUBLIC_KEY            = "test_instance_public_key"
)

func buildTestInstanceJsonResponse(instance *Instance) []byte {
	return []byte(`{"id": "` + instance.Id + `", ` +
		`"name":"` + instance.Name + `", ` +
		`"state":"` + instance.State + `", ` +
		`"templateId":"` + instance.TemplateId + `", ` +
		`"templateName":"` + instance.TemplateName + `", ` +
		`"isPasswordEnabled":` + strconv.FormatBool(instance.IsPasswordEnabled) + `, ` +
		`"isSshKeyEnabled":` + strconv.FormatBool(instance.IsSSHKeyEnabled) + `, ` +
		`"username":"` + instance.Username + `", ` +
		`"computeOfferingId":"` + instance.ComputeOfferingId + `", ` +
		`"computeOfferingName":"` + instance.ComputeOfferingName + `", ` +
		`"cpuCount": ` + strconv.Itoa(instance.CpuCount) + `, ` +
		`"memoryInMB": ` + strconv.Itoa(instance.MemoryInMB) + `, ` +
		`"zoneId":"` + instance.ZoneId + `", ` +
		`"zoneName":"` + instance.ZoneName + `", ` +
		`"projectId":"` + instance.ProjectId + `", ` +
		`"networkId":"` + instance.NetworkId + `", ` +
		`"networkName":"` + instance.NetworkName + `", ` +
		`"vpcId":"` + instance.VpcId + `", ` +
		`"vpcName":"` + instance.VpcName + `", ` +
		`"macAddress":"` + instance.MacAddress + `", ` +
		`"ipAddress":"` + instance.IpAddress + `", ` +
		`"volumeIdToAttach":"` + instance.VolumeIdToAttach + `", ` +
		`"publicKey":"` + instance.PublicKey + `", ` +
		`"userData":"` + instance.UserData + `"}`)
}

func buildListTestInstanceJsonResponse(instances []Instance) []byte {
	resp := `[`
	for i, inst := range instances {
		resp += string(buildTestInstanceJsonResponse(&inst))
		if i != len(instances)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetInstanceReturnInstanceIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	expectedInstance := Instance{Id: TEST_INSTANCE_ID,
		Name:                TEST_INSTANCE_NAME,
		State:               TEST_INSTANCE_STATE,
		TemplateId:          TEST_INSTANCE_TEMPLATE_ID,
		TemplateName:        TEST_INSTANCE_TEMPLATE_NAME,
		IsPasswordEnabled:   TEST_INSTANCE_IS_PASSWORD_ENABLED,
		IsSSHKeyEnabled:     TEST_INSTANCE_IS_SSH_KEY_ENABLED,
		Username:            TEST_INSTANCE_USERNAME,
		ComputeOfferingId:   TEST_INSTANCE_COMPUTE_OFFERING_ID,
		ComputeOfferingName: TEST_INSTANCE_COMPUTE_OFFERING_NAME,
		CpuCount:            TEST_INSTANCE_CPU_NUMBER,
		MemoryInMB:          TEST_INSTANCE_MEMORY_IN_MB,
		ZoneId:              TEST_INSTANCE_ZONE_ID,
		ZoneName:            TEST_INSTANCE_ZONE_NAME,
		ProjectId:           TEST_INSTANCE_PROJECT_ID,
		NetworkId:           TEST_INSTANCE_NETWORK_ID,
		NetworkName:         TEST_INSTANCE_NETWORK_NAME,
		VpcId:               TEST_INSTANCE_VPC_ID,
		VpcName:             TEST_INSTANCE_VPC_NAME,
		MacAddress:          TEST_INSTANCE_MAC_ADDRESS,
		IpAddress:           TEST_INSTANCE_IP_ADDRESS,
		VolumeIdToAttach:    TEST_INSTANCE_VOLUME_ID_TO_ATTACH,
		PublicKey:           TEST_INSTANCE_PUBLIC_KEY,
		UserData:            TEST_INSTANCE_USER_DATA}

	mockEntityService.EXPECT().Get(TEST_INSTANCE_ID, gomock.Any()).Return(buildTestInstanceJsonResponse(&expectedInstance), nil)

	//when
	instance, _ := instanceService.Get(TEST_INSTANCE_ID)

	//then
	if assert.NotNil(t, instance) {
		assert.Equal(t, expectedInstance, *instance)
	}
}

func TestGetInstanceReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_INSTANCE_ID, gomock.Any()).Return(nil, mockError)

	//when
	instance, err := instanceService.Get(TEST_INSTANCE_ID)

	//then
	assert.Nil(t, instance)
	assert.Equal(t, mockError, err)

}

func TestListInstanceReturnInstancesIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	expectedInstance1 := Instance{Id: "list_id_1",
		Name:                "list_name_1",
		State:               "list_state_1",
		TemplateId:          "list_template_id_1",
		TemplateName:        "list_template_name_1",
		IsPasswordEnabled:   false,
		IsSSHKeyEnabled:     true,
		Username:            "list_username_1",
		ComputeOfferingId:   "list_compute_offering_id_1",
		ComputeOfferingName: "list_compute_offering_name_1",
		CpuCount:            2,
		MemoryInMB:          12425,
		ZoneId:              "list_zone_id_1",
		ZoneName:            "list_zone_name_1",
		ProjectId:           "list_project_id_1",
		NetworkId:           "list_network_id_1",
		NetworkName:         "list_network_name_1",
		VpcId:               "list_vpc_id_1",
		VpcName:             "list_vpc_name_1",
		MacAddress:          "list_mac_address_1",
		VolumeIdToAttach:    "list_volume_id_to_attach_1",
		IpAddress:           "list_ip_address_1",
		PublicKey:           "list_public_key_1",
		UserData:            "list_user_data_1"}

	expectedInstances := []Instance{expectedInstance1}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestInstanceJsonResponse(expectedInstances), nil)

	//when
	instances, _ := instanceService.List()

	//then
	if assert.NotNil(t, instances) {
		assert.Equal(t, expectedInstances, instances)
	}
}

func TestListInstanceReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	instances, err := instanceService.List()

	//then
	assert.Nil(t, instances)
	assert.Equal(t, mockError, err)

}

func TestCreateInstanceReturnCreatedInstanceIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	instanceToCreate := Instance{Id: "new_id",
		Name:              "new_name",
		TemplateId:        "templateId",
		ComputeOfferingId: "computeOfferingId",
		NetworkId:         "networkId"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{"id":"new_id", "password": "new_password"}`), nil)

	//when
	createdInstance, _ := instanceService.Create(instanceToCreate)

	//then
	if assert.NotNil(t, createdInstance) {
		assert.Equal(t, "new_password", createdInstance.Password)
	}
}

func TestCreateInstanceReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_instance_error"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	instanceToCreate := Instance{Name: "new_name",
		TemplateId:        "templateId",
		ComputeOfferingId: "computeOfferingId",
		NetworkId:         "networkId"}

	//when
	createdInstance, err := instanceService.Create(instanceToCreate)

	//then
	assert.Nil(t, createdInstance)
	assert.Equal(t, mockError, err)

}

func TestPurgeInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_PURGE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.Purge(TEST_INSTANCE_ID)

	//then
	assert.True(t, success)
}

func TestPurgeInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_purge_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_PURGE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.Purge(TEST_INSTANCE_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestStartInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_START_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.Start(TEST_INSTANCE_ID)

	//then
	assert.True(t, success)
}

func TestStartInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_start_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_START_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.Start(TEST_INSTANCE_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestStopInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_STOP_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.Stop(TEST_INSTANCE_ID)

	//then
	assert.True(t, success)
}

func TestStopInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_stop_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_STOP_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.Stop(TEST_INSTANCE_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestDestroyInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Delete(TEST_INSTANCE_ID, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.Destroy(TEST_INSTANCE_ID, false)

	//then
	assert.True(t, success)
}

func TestDestroyInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_destroy_instance_error"}
	mockEntityService.EXPECT().Delete(TEST_INSTANCE_ID, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.Destroy(TEST_INSTANCE_ID, true)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestDestroyWithOptionsInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Delete(TEST_INSTANCE_ID, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.DestroyWithOptions(TEST_INSTANCE_ID, DestroyOptions{DeleteSnapshots: true})

	//then
	assert.True(t, success)
}

func TestDestroyWithOptionsInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_destroy_instance_error"}
	mockEntityService.EXPECT().Delete(TEST_INSTANCE_ID, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.DestroyWithOptions(TEST_INSTANCE_ID, DestroyOptions{PurgeImmediately: true})

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestRecoverInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_RECOVER_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.Recover(TEST_INSTANCE_ID)

	//then
	assert.True(t, success)
}

func TestRecoverInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_recover_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_RECOVER_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.Recover(TEST_INSTANCE_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestRebootInstanceReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_REBOOT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.Reboot(TEST_INSTANCE_ID)

	//then
	assert.True(t, success)
}

func TestRebootInstanceReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_reboot_instance_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_REBOOT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.Reboot(TEST_INSTANCE_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestAssociateSSHKeyReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_ASSOCIATE_SSH_KEY_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.AssociateSSHKey(TEST_INSTANCE_ID, "new_ssh_key")

	//then
	assert.True(t, success)
}

func TestAssociateSSHKeyReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_associate_ssh_key_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_ASSOCIATE_SSH_KEY_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.AssociateSSHKey(TEST_INSTANCE_ID, "new_ssh_key")

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestChangeComputeOfferingReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	instanceWithNewComputeOffering := Instance{
		Id:                TEST_INSTANCE_ID,
		ComputeOfferingId: "new_compute_offering",
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_CHANGE_COMPUTE_OFFERING_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.ChangeComputeOffering(instanceWithNewComputeOffering)

	//then
	assert.True(t, success)
}

func TestChangeComputeOfferingReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_change_compute_offering_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_CHANGE_COMPUTE_OFFERING_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	instanceWithNewComputeOffering := Instance{
		Id:                TEST_INSTANCE_ID,
		ComputeOfferingId: "new_compute_offering",
	}

	//when
	success, err := instanceService.ChangeComputeOffering(instanceWithNewComputeOffering)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestResetPasswordReturnNewPasswordIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_RESET_PASSWORD_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{"password":"new_password"}`), nil)

	//when
	newPassword, _ := instanceService.ResetPassword(TEST_INSTANCE_ID)

	//then
	assert.Equal(t, "new_password", newPassword)
}

func TestResetPasswordReturnEmptyStringIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_reset_password_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_RESET_PASSWORD_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	newPassword, err := instanceService.ResetPassword(TEST_INSTANCE_ID)

	//then
	assert.Empty(t, newPassword)
	assert.Equal(t, mockError, err)
}

func TestCreateRecoveryPointReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_CREATE_RECOVERY_POINT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := instanceService.CreateRecoveryPoint(TEST_INSTANCE_ID, RecoveryPoint{"new_recovery_point_name", "description"})

	//then
	assert.True(t, success)
}

func TestCreateRecoveryPointReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_recovery_point_error"}
	mockEntityService.EXPECT().Execute(TEST_INSTANCE_ID, INSTANCE_CREATE_RECOVERY_POINT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := instanceService.CreateRecoveryPoint(TEST_INSTANCE_ID, RecoveryPoint{"new_recovery_point_name", "description"})

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestExistsReturnTrueIfInstanceExists(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Get(TEST_INSTANCE_ID, gomock.Any()).Return([]byte(`{"id": "foo"}`), nil)

	//when
	exists, _ := instanceService.Exists(TEST_INSTANCE_ID)

	//then
	assert.True(t, exists)
}

func TestExistsReturnFalseIfInstanceDoesntExist(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockApiError := api.CcaErrorResponse(api.CcaResponse{StatusCode: api.NOT_FOUND})
	mockEntityService.EXPECT().Get(TEST_INSTANCE_ID, gomock.Any()).Return([]byte(`{}`), mockApiError)

	//when
	exists, err := instanceService.Exists(TEST_INSTANCE_ID)

	//then
	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestExistsReturnErrorIfUnexpectedError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	instanceService := InstanceApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_exists_error"}
	mockEntityService.EXPECT().Get(TEST_INSTANCE_ID, gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	_, err := instanceService.Exists(TEST_INSTANCE_ID)

	//then
	assert.Equal(t, mockError, err)
}
