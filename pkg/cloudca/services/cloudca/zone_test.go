package cloudca

import (
	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetZoneByIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert := assert.New(t)

	// given
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	zoneService := ZoneApi{
		entityService: mockEntityService,
	}
	zoneId := "zoneid"
	mockEntityService.EXPECT().Get(zoneId, gomock.Any()).Return([]byte(`{"id":"zoneid","name":"zonename"}`), nil)

	// when
	zone, err := zoneService.Get(zoneId)

	// then
	assert.Nil(err)
	assert.Equal(zoneId, zone.Id)
}

func TestGetZoneByIdFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert := assert.New(t)

	// given
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	zoneService := ZoneApi{
		entityService: mockEntityService,
	}
	zoneId := "zoneid"
	mockError := mocks.MockError{"fetch error"}
	mockEntityService.EXPECT().Get(zoneId, gomock.Any()).Return(nil, mockError)

	// when
	zone, err := zoneService.Get(zoneId)

	// then
	assert.Nil(zone)
	assert.Equal(mockError, err)
}

func TestListZonesSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert := assert.New(t)

	// given
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	zoneService := ZoneApi{
		entityService: mockEntityService,
	}
	zone1 := Zone{
		Id:   "zoneid1",
		Name: "zone1",
	}
	zone2 := Zone{
		Id:   "zoneid2",
		Name: "zone2",
	}
	allZones := []Zone{zone1, zone2}
	mockEntityService.EXPECT().List(gomock.Any()).Return([]byte(`[{"id":"zoneid1","name":"zone1"},{"id":"zoneid2","name":"zone2"}]`), nil)

	// when
	zones, err := zoneService.List()

	// then
	assert.Nil(err)
	assert.Equal(zones, allZones)
}

func TestListZonesFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert := assert.New(t)

	// given
	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	zoneService := ZoneApi{
		entityService: mockEntityService,
	}
	mockError := mocks.MockError{"list error"}
	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	// when
	zones, err := zoneService.List()

	// then
	assert.Nil(zones)
	assert.Equal(mockError, err)
}
