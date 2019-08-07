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
	TEST_VOLUME_ID               = "test_volume_id"
	TEST_VOLUME_NAME             = "test_volume"
	TEST_VOLUME_TYPE             = "test_volume_type"
	TEST_VOLUME_CREATION_DATE    = "test_volume_creation_date"
	TEST_VOLUME_SIZE             = 500
	TEST_VOLUME_DISK_OFFERING_ID = "test_volume_disk_offering_id"
	TEST_VOLUME_TEMPLATE_ID      = "test_volume_template_id"
	TEST_VOLUME_ZONE_NAME        = "test_volume_zone_name"
	TEST_VOLUME_STATE            = "test_volume_state"
	TEST_VOLUME_INSTANCE_NAME    = "test_volume_instance_name"
	TEST_VOLUME_INSTANCE_ID      = "test_volume_instance_id"
	TEST_VOLUME_INSTANCE_STATE   = "test_volume_instance_state"
)

func buildVolumeJsonResponse(volume *Volume) []byte {
	return []byte(`{"id":"` + volume.Id + `",` +
		` "name": "` + volume.Name + `",` +
		` "type": "` + volume.Type + `",` +
		` "creationDate": "` + volume.CreationDate + `",` +
		` "size": ` + strconv.Itoa(volume.Size) + `,` +
		` "sizeInGb": ` + strconv.Itoa(volume.GbSize) + `,` +
		` "diskOfferingId": "` + volume.DiskOfferingId + `",` +
		` "templateId": "` + volume.TemplateId + `",` +
		` "zoneName": "` + volume.ZoneName + `",` +
		` "state": "` + volume.State + `",` +
		` "instanceName": "` + volume.InstanceName + `",` +
		` "instanceId": "` + volume.InstanceId + `",` +
		` "instanceState": "` + volume.InstanceState + `"}`)
}

func buildListVolumeJsonResponse(volumes []Volume) []byte {
	resp := `[`
	for i, v := range volumes {
		resp += string(buildVolumeJsonResponse(&v))
		if i != len(volumes)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetVolumeReturnVolumeIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	expectedVolume := Volume{Id: TEST_VOLUME_ID,
		Name:           TEST_VOLUME_NAME,
		Type:           TEST_VOLUME_TYPE,
		CreationDate:   TEST_VOLUME_CREATION_DATE,
		Size:           TEST_VOLUME_SIZE,
		DiskOfferingId: TEST_VOLUME_DISK_OFFERING_ID,
		TemplateId:     TEST_VOLUME_TEMPLATE_ID,
		ZoneName:       TEST_VOLUME_ZONE_NAME,
		State:          TEST_VOLUME_STATE,
		InstanceName:   TEST_VOLUME_INSTANCE_NAME,
		InstanceId:     TEST_VOLUME_INSTANCE_ID,
		InstanceState:  TEST_VOLUME_INSTANCE_STATE}

	mockEntityService.EXPECT().Get(TEST_VOLUME_ID, gomock.Any()).Return(buildVolumeJsonResponse(&expectedVolume), nil)

	//when
	volume, _ := volumeService.Get(TEST_VOLUME_ID)

	//then
	if assert.NotNil(t, volume) {
		assert.Equal(t, expectedVolume, *volume)
	}
}

func TestGetVolumeReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_VOLUME_ID, gomock.Any()).Return(nil, mockError)

	//when
	volume, err := volumeService.Get(TEST_VOLUME_ID)

	//then
	assert.Nil(t, volume)
	assert.Equal(t, mockError, err)

}

func TestListVolumeReturnVolumesIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	expectedVolume1 := Volume{Id: "list_id_1",
		Name:           "list_name_1",
		Type:           "list_type_1",
		CreationDate:   "list_creation_date_1",
		Size:           1215,
		DiskOfferingId: "list_disk_offering_id_1",
		TemplateId:     "list_template_id_1",
		ZoneName:       "list_zone_name_1",
		State:          "list_state_1",
		InstanceName:   "list_instance_name_1",
		InstanceId:     "list_instance_id_1",
		InstanceState:  "list_instance_state_1"}
	expectedVolume2 := Volume{Id: "list_id_2",
		Name:           "list_name_2",
		Type:           "list_type_2",
		CreationDate:   "list_creation_date_2",
		Size:           54582,
		DiskOfferingId: "list_disk_offering_id_2",
		TemplateId:     "list_template_id_2",
		ZoneName:       "list_zone_name_2",
		State:          "list_state_2",
		InstanceName:   "list_instance_name_2",
		InstanceId:     "list_instance_id_2",
		InstanceState:  "list_instance_state_2"}

	expectedVolumes := []Volume{expectedVolume1, expectedVolume2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListVolumeJsonResponse(expectedVolumes), nil)

	//when
	volumes, _ := volumeService.List()

	//then
	if assert.NotNil(t, volumes) {
		assert.Equal(t, expectedVolumes, volumes)
	}
}

func TestListVolumeReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	volumes, err := volumeService.List()

	//then
	assert.Nil(t, volumes)
	assert.Equal(t, mockError, err)

}

func TestCreateReturnsErrorIfErrorWhileCreating(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	mockError := mocks.MockError{"creation error"}
	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	volume, err := volumeService.Create(Volume{})

	// then
	assert.Nil(t, volume)
	assert.Equal(t, mockError, err)
}

func TestCreateSucceeds(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	expected := Volume{
		Id: "expected",
	}
	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(buildVolumeJsonResponse(&expected), nil)

	// when
	volume, err := volumeService.Create(Volume{})

	// then
	assert.Equal(t, expected, *volume)
	assert.Nil(t, err)
}

func TestAttachToInstanceFailure(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	toAttach := &Volume{
		Id: "toAttach",
	}
	mockError := mocks.MockError{"attach error"}
	mockEntityService.EXPECT().Execute(toAttach.Id, "attachToInstance", gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	err := volumeService.AttachToInstance(toAttach, "some instance")

	// then
	assert.Equal(t, mockError, err)
}

func TestAttachToInstanceSuccess(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	toAttach := &Volume{
		Id: "toAttach",
	}
	mockEntityService.EXPECT().Execute(toAttach.Id, "attachToInstance", gomock.Any(), gomock.Any()).Return([]byte("success!"), nil)

	// when
	err := volumeService.AttachToInstance(toAttach, "some instance")

	// then
	assert.Nil(t, err)
}

func TestDetachFromInstanceFailure(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	toDetach := &Volume{
		Id: "toDetach",
	}
	mockError := mocks.MockError{"Detach error"}
	mockEntityService.EXPECT().Execute(toDetach.Id, "detachFromInstance", gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	err := volumeService.DetachFromInstance(toDetach)

	// then
	assert.Equal(t, mockError, err)
}

func TestDetachFromInstanceSuccess(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	toDetach := &Volume{
		Id: "toDetach",
	}
	mockEntityService.EXPECT().Execute(toDetach.Id, "detachFromInstance", gomock.Any(), gomock.Any()).Return([]byte("success!"), nil)

	// when
	err := volumeService.DetachFromInstance(toDetach)

	// then
	assert.Nil(t, err)
}

func TestDeleteSuccess(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	toDelete := &Volume{
		Id: "toDelete",
	}
	mockEntityService.EXPECT().Delete(toDelete.Id, gomock.Any(), gomock.Any()).Return([]byte("success!"), nil)

	// when
	err := volumeService.Delete(toDelete.Id)

	// then
	assert.Nil(t, err)
}

func TestDeleteFailure(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	volumeService := VolumeApi{
		entityService: mockEntityService,
	}
	toDelete := &Volume{
		Id: "toDelete",
	}
	mockError := mocks.MockError{"delete error"}
	mockEntityService.EXPECT().Delete(toDelete.Id, gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	err := volumeService.Delete(toDelete.Id)

	// then
	assert.Equal(t, mockError, err)
}
