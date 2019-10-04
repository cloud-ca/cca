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

// Package cli wraps around and holds the references to different part of cca cli command
package cli

import (
	"github.com/cloud-ca/cca/pkg/cli/client"
	"github.com/cloud-ca/cca/pkg/cli/flags"
	"github.com/cloud-ca/cca/pkg/cli/output"
)

// Wrapper of different parts of cca cli
type Wrapper struct {
	GlobalFlags   *flags.GlobalFlags
	OutputBuilder *output.Builder
	CcaClient     *client.Client
}
