// Copyright © 2019 cloud.ca Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cloudca contains the implementation of cloud.ca API
package cloudca

import (
	"github.com/cloud-ca/cca/pkg/cloudca/api"
	"github.com/cloud-ca/cca/pkg/cloudca/configuration"
	"github.com/cloud-ca/cca/pkg/cloudca/services"
	"github.com/cloud-ca/cca/pkg/cloudca/services/cloudca"
)

// DefaultAPIURL default API URL to use
const DefaultAPIURL = "https://api.cloud.ca/v1/"

// Client for interacting with cloud.ca API
type Client struct {
	apiClient          api.Client
	Tasks              services.TaskService
	Environments       configuration.EnvironmentService
	Users              configuration.UserService
	ServiceConnections configuration.ServiceConnectionService
	Organizations      configuration.OrganizationService
}

// GetAPIClient Get the API Client used by all the services
func (c Client) GetAPIClient() api.Client {
	return c.apiClient
}

// GetAPIURL Get the API url used to do he calls
func (c Client) GetAPIURL() string {
	return c.GetAPIClient().GetAPIURL()
}

// GetAPIKey Get the API key used in the calls
func (c Client) GetAPIKey() string {
	return c.GetAPIClient().GetAPIKey()
}

// GetResources get the Resources for a specific serviceCode and environmentName
// For now it assumes that the serviceCode belongs to a cloud.ca service type
func (c Client) GetResources(serviceCode string, environmentName string) (services.ServiceResources, error) {
	//TODO: change to check service type of service code
	return cloudca.NewResources(c.apiClient, serviceCode, environmentName), nil
}

// NewClient Create a Client with the default URL
func NewClient(apiKey string) *Client {
	return NewClientWithURL(DefaultAPIURL, apiKey)
}

// NewClientWithURL Create a Client with a custom URL
func NewClientWithURL(apiURL string, apiKey string) *Client {
	apiClient := api.NewClient(apiURL, apiKey)
	return NewClientWithAPIClient(apiClient)
}

// NewInsecureClientWithURL Create a Client with a custom URL that accepts insecure connections
func NewInsecureClientWithURL(apiURL string, apiKey string) *Client {
	apiClient := api.NewInsecureClient(apiURL, apiKey)
	return NewClientWithAPIClient(apiClient)
}

// NewClientWithAPIClient Create a Client with a provided API client
func NewClientWithAPIClient(apiClient api.Client) *Client {
	return &Client{
		apiClient:          apiClient,
		Tasks:              services.NewTaskService(apiClient),
		Environments:       configuration.NewEnvironmentService(apiClient),
		Users:              configuration.NewUserService(apiClient),
		ServiceConnections: configuration.NewServiceConnectionService(apiClient),
		Organizations:      configuration.NewOrganizationService(apiClient),
	}
}

///////////////////////////////////////////////////
// Deperacated functions
///////////////////////////////////////////////////

// NewCcaClient Create a Client with the default URL
// **Deprecated, Use NewClient(apiKey string) instead**
func NewCcaClient(apiKey string) *Client {
	return NewClient(apiKey)
}

// NewCcaClientWithURL Create a Client with a custom URL
// **Deprecated, NewClientWithURL(apiKey string, apiKey string) instead**
func NewCcaClientWithURL(apiURL string, apiKey string) *Client {
	return NewClientWithURL(apiURL, apiKey)
}

// NewInsecureCcaClientWithURL Create a Client with a custom URL that accepts insecure connections
// **Deprecated, Use NewInsecureClientWithURL(apiURL string, apiKey string) instead**
func NewInsecureCcaClientWithURL(apiURL string, apiKey string) *Client {
	return NewInsecureClientWithURL(apiURL, apiKey)
}

// NewCcaClientWithAPIClient Create a Client with a provided API client
// **Deprecated, Use NewCcaClientWithApiClient(apiClient api.Client) instead**
func NewCcaClientWithAPIClient(apiClient api.Client) *Client {
	return NewClientWithAPIClient(apiClient)
}
