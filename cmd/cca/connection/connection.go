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

// Package connection implements the `connection` command
package connection

import (
	"github.com/cloud-ca/cca/pkg/cli"
	"github.com/cloud-ca/cca/pkg/util"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for connection
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "connection",
		Short: "Manage service connections that you can create resources for",
		Long: util.LongDescription(`
            Service connections are the services that you can create resources for (e.g. compute, object
            storage). Environments are created for a specific service which allows you to create and
            manage resources within that service.
        `),
	}

	return cmd
}
