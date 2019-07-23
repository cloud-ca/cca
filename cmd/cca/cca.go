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

package cca

import (
	"os"

	"github.com/cloud-ca/cca/cmd/cca/completion"
	"github.com/cloud-ca/cca/cmd/cca/version"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command implementing the root command for cca
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:         cobra.NoArgs,
		Use:          "cca",
		Short:        "cca CLI manages authentication, configurations and interactions with the cloud.ca APIs.",
		Long:         "cca CLI manages authentication, configurations and interactions with the cloud.ca APIs.",
		SilenceUsage: true,
		Version:      version.Version(),
	}

	// add all top level subcommands
	cmd.AddCommand(completion.NewCommand())
	cmd.AddCommand(version.NewCommand())

	return cmd
}

// Run runs the `cca` root command
func Run() error {
	return NewCommand().Execute()
}

// Main wraps Run
func Main() {
	if err := Run(); err != nil {
		os.Exit(1)
	}
}
