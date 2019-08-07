package cloudca

import (
	"fmt"
	"testing"

	"github.com/cloud-ca/cca/pkg/cloudca/mocks"
	"github.com/cloud-ca/cca/pkg/cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const ACL_RULE_TEMPLATE = `{
	"id": "%s",
	"networkAclId": "6145ea41-010c-41f2-a065-2a3a4e98d09d",
	"ruleNumber": "1",
	"cidr": "0.0.0.0/24",
	"action": "Allow",
	"protocol": "TCP",
	"startPort": "80",
	"endPort": "80",
	"trafficType": "Ingress",
	"state": "Active"
}`

func setupMockForNetworkAclRule(t *testing.T) *services_mocks.MockEntityService {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return services_mocks.NewMockEntityService(ctrl)
}

func createNetworkAclRuleWithId(id string) *NetworkAclRule {
	return &NetworkAclRule{
		Id:           id,
		NetworkAclId: "6145ea41-010c-41f2-a065-2a3a4e98d09d",
		RuleNumber:   "1",
		Cidr:         "0.0.0.0/24",
		Action:       "Allow",
		Protocol:     "TCP",
		StartPort:    "80",
		EndPort:      "80",
		TrafficType:  "Ingress",
		State:        "Active",
	}
}

func buildTestNetworkAclRuleJsonResponse(networkAclRule *NetworkAclRule) []byte {
	return []byte(`{"id":"` + networkAclRule.Id + `",` +
		` "networkAclId":"` + networkAclRule.NetworkAclId + `",` +
		` "ruleNumber":"` + networkAclRule.RuleNumber + `",` +
		` "cidr":"` + networkAclRule.Cidr + `",` +
		` "action":"` + networkAclRule.Action + `",` +
		` "protocol":"` + networkAclRule.Protocol + `",` +
		` "startPort":"` + networkAclRule.StartPort + `",` +
		` "endPort":"` + networkAclRule.EndPort + `",` +
		` "trafficType":"` + networkAclRule.TrafficType + `",` +
		` "state":"` + networkAclRule.State +
		`"}`)
}

func buildListTestNetworkAclRulesJsonResponse(acls []NetworkAclRule) []byte {
	resp := `[`
	for i, t := range acls {
		resp += string(buildTestNetworkAclRuleJsonResponse(&t))
		if i != len(acls)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetNetworkAclRuleByIdReturnAclRule_ifSuccess(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "rule_0"
	expectedNetworkAclRule := *createNetworkAclRuleWithId(expectedId)

	response := fmt.Sprintf(ACL_RULE_TEMPLATE, expectedId)
	mockEntityService.EXPECT().Get(expectedId, gomock.Any()).Return([]byte(response), nil)

	// when
	networkAclRule, _ := networkAclRuleService.Get(expectedId)

	// then
	assert.Equal(t, expectedNetworkAclRule, *networkAclRule)
}

func TestGetNetworkAclRuleByIdReturnError_ifError(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "rule_0"
	mockError := mocks.MockError{Message: "get error"}
	mockEntityService.EXPECT().Get(expectedId, gomock.Any()).Return(nil, mockError)

	// when
	networkAclRule, err := networkAclRuleService.Get(expectedId)

	// then
	assert.Nil(t, networkAclRule)
	assert.Equal(t, mockError, err)
}

func TestListNetworkAclRuleReturnAclsIfSuccess(t *testing.T) {
	//given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId1 := "rule_1"
	expectedNetworkAclRule1 := *createNetworkAclRuleWithId(expectedId1)
	expectedId2 := "rule_2"
	expectedNetworkAclRule2 := *createNetworkAclRuleWithId(expectedId2)

	expectedAcls := []NetworkAclRule{expectedNetworkAclRule1, expectedNetworkAclRule2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestNetworkAclRulesJsonResponse(expectedAcls), nil)

	//when
	acls, _ := networkAclRuleService.List()

	//then
	if assert.NotNil(t, acls) {
		assert.Equal(t, expectedAcls, acls)
	}
}

func TestListNetworkAclRuleReturnNilWithErrorIfError(t *testing.T) {
	//given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{Message: "some_list_error"}
	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	acls, err := networkAclRuleService.List()

	//then
	assert.Nil(t, acls)
	assert.Equal(t, mockError, err)

}

func TestListNetworkAclRuleByAclIdReturnAcls_ifSuccess(t *testing.T) {
	//given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId1 := "rule_1"
	expectedNetworkAclRule1 := *createNetworkAclRuleWithId(expectedId1)
	expectedAcls := []NetworkAclRule{expectedNetworkAclRule1}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestNetworkAclRulesJsonResponse(expectedAcls), nil)

	//when
	acls, _ := networkAclRuleService.ListByNetworkAclId("acl1")

	//then
	if assert.NotNil(t, acls) {
		assert.Equal(t, expectedAcls, acls)
	}
}

func TestListNetworkAclRuleByAclIdReturnNilWithError_ifError(t *testing.T) {
	//given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{Message: "some_list_error"}
	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	acls, err := networkAclRuleService.ListByNetworkAclId("acl1")

	//then
	assert.Nil(t, acls)
	assert.Equal(t, mockError, err)

}

func TestListNetworkAclRulesWithOptions(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	id1, id2 := "1234", "4321"
	rule1, rule2 := fmt.Sprintf(ACL_RULE_TEMPLATE, id1), fmt.Sprintf(ACL_RULE_TEMPLATE, id2)
	response := fmt.Sprintf("[ %s, %s ]", rule1, rule2)
	mockEntityService.EXPECT().List(gomock.Any()).Return([]byte(response), nil)

	// when
	rules, _ := networkAclRuleService.ListWithOptions(map[string]string{})

	// then
	assert.Equal(t, id1, rules[0].Id)
	assert.Equal(t, id2, rules[1].Id)
}

func TestCreateNetworkAclRuleReturnsError_ifErrorWhileCreating(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{Message: "creation error"}
	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	rule, err := networkAclRuleService.Create(NetworkAclRule{})

	// then
	assert.Nil(t, rule)
	assert.Equal(t, mockError, err)
}

func TestCreateNetworkAclRuleReturnsSuccess_ifNoErrorsOccur(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "adsf"
	response := fmt.Sprintf(ACL_RULE_TEMPLATE, expectedId)
	expectedNetworkAclRule := *createNetworkAclRuleWithId(expectedId)

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(response), nil)

	// when
	rule, _ := networkAclRuleService.Create(expectedNetworkAclRule)

	// then
	assert.Equal(t, expectedNetworkAclRule, *rule)
}

func TestUpdateNetworkAclRuleReturnsError_ifErrorWhileCreating(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{Message: "update error"}
	mockEntityService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, mockError)

	// when
	rule, err := networkAclRuleService.Update("1234", NetworkAclRule{})

	// then
	assert.Nil(t, rule)
	assert.Equal(t, mockError, err)
}

func TestUpdateNetworkAclRuleReturnsSuccess_ifNoErrorsOccur(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "adsf"
	response := fmt.Sprintf(ACL_RULE_TEMPLATE, expectedId)
	expectedNetworkAclRule := *createNetworkAclRuleWithId(expectedId)

	mockEntityService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(response), nil)

	// when
	rule, _ := networkAclRuleService.Update(expectedId, expectedNetworkAclRule)

	// then
	assert.Equal(t, expectedNetworkAclRule, *rule)
}

func TestDeleteNetworkAclRuleReturnsError_ifErrorWhileDeleting(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{Message: "deletion error"}
	mockEntityService.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte{}, mockError)

	// when
	success, err := networkAclRuleService.Delete("123")

	// then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestDeleteNetworkAclRuleReturnsSuccess_ifNoErrorsOccur(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "id0"
	mockEntityService.EXPECT().Delete(expectedId, gomock.Any(), gomock.Any()).Return([]byte{}, nil)

	// when
	success, _ := networkAclRuleService.Delete(expectedId)

	// then
	assert.True(t, success)
}
