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

// Package list implements the `connection list` command
package list

import (
	"github.com/cloud-ca/cca/pkg/cli"
	"github.com/cloud-ca/cca/pkg/output"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for connection list
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		Use:     "list",
		Short:   "List all service connections",
		Long:    "List all service connections",
		RunE: func(cmd *cobra.Command, args []string) error {
			connections, err := cli.CcaClient.ServiceConnections.List()
			if err != nil {
				return err
			}
			return cli.OutputBuilder.Build(func(formatter *output.Formatter) error {
				return formatter.Format(connections)
			})
		},
	}

	return cmd
}
