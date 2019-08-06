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

// Package environment implements the `environment` command
package environment

import (
	"github.com/cloud-ca/cca/cmd/cca/environment/delete"
	"github.com/cloud-ca/cca/cmd/cca/environment/list"
	"github.com/cloud-ca/cca/pkg/cli"
	"github.com/cloud-ca/cca/pkg/util"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for environment actions
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Aliases: []string{"env"},
		Use:     "environment",
		Short:   "Manage resources of a specific service and users’ access to them",
		Long: util.LongDescription(`
            Environments allow you to manage resources of a specific service and to manage your users’
            access to them. With environment roles, you have tight control of what a user is allowed to
            do in your environment. A general use case of environments is to split your resources into
            different deployment environments (e.g. dev, staging and production). The advantage is that
            resources of different deployments are isolated from each other and you can restrict user
            access to your most critical resources.
        `),
	}

	cmd.AddCommand(delete.NewCommand(cli))
	cmd.AddCommand(list.NewCommand(cli))

	return cmd
}
