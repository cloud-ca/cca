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

// Package list implements the `environment list` command
package list

import (
	"github.com/cloud-ca/cca/pkg/cli"
	"github.com/cloud-ca/cca/pkg/output"
	"github.com/cloud-ca/cca/pkg/util"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for environment list
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		Use:     "list",
		Short:   "List environments that you have access to",
		Long: util.LongDescription(`
            List environments that you have access to. It will only return environments that you’re
            member of if you’re not assigned the Environments read permission.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			envs, err := cli.CcaClient.Environments.List()
			if err != nil {
				return err
			}
			return cli.OutputBuilder.Build(func(formatter *output.Formatter) error {
				return formatter.Format(envs)
			})
		},
	}

	return cmd
}
