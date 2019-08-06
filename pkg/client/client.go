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

// Package client contains the client library to interact with cloud.ca infrastructure
package client

import (
	gocca "github.com/cloud-ca/go-cloudca"
)

// Client to interact with cloud.ca infrastructure
type Client struct {
	*gocca.CcaClient
}

// NewClient returns a new client to interact with cloud.ca
// infrastructure with provided API URL and Key
func NewClient(url string, key string) *Client {
	return &Client{
		CcaClient: gocca.NewCcaClientWithURL(url, key),
	}
}
