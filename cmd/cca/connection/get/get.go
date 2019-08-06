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

// Package get implements the `connection get` command
package get

import (
	"github.com/cloud-ca/cca/pkg/cli"
	"github.com/cloud-ca/cca/pkg/output"
	"github.com/spf13/cobra"
)

type flag struct {
	id string
}

// NewCommand returns a new cobra.Command for connection get
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	flg := &flag{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "get",
		Short: "Get service connection",
		Long:  "Get service connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			connection, err := cli.CcaClient.ServiceConnections.Get(flg.id)
			if err != nil {
				return err
			}
			return cli.OutputBuilder.Build(func(formatter *output.Formatter) error {
				return formatter.Format(connection)
			})
		},
	}

	cmd.Flags().StringVar(&flg.id, "id", "", "environment id")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}

	return cmd
}
