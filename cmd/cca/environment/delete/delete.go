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

// Package delete implements the `environment delete` command
package delete

import (
	"github.com/cloud-ca/cca/pkg/cli"
	"github.com/cloud-ca/cca/pkg/output"
	"github.com/cloud-ca/cca/pkg/util"
	"github.com/spf13/cobra"
)

type flag struct {
	id string
}

// NewCommand returns a new cobra.Command for environment delete
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	flg := &flag{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "delete",
		Short: "Delete a specific environment",
		Long: util.LongDescription(`
            Delete a specific environment. You will need a role with the Delete an existing environment
            permission to execute this operation.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			deleted, err := cli.CcaClient.Environments.Delete(flg.id)
			if err != nil {
				return err
			}
			return cli.OutputBuilder.Build(func(formatter *output.Formatter) error {
				type R struct {
					Deleted bool `json:"deleted"`
				}
				return formatter.Format(&R{Deleted: deleted})
			})
		},
	}

	cmd.Flags().StringVar(&flg.id, "id", "", "ID of environment to delete")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}

	return cmd
}
