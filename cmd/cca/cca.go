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

// Package cca implements root command
package cca

import (
	"os"

	"github.com/cloud-ca/cca/cmd/cca/completion"
	"github.com/cloud-ca/cca/cmd/cca/version"
	"github.com/cloud-ca/cca/pkg/cmdutil"
	"github.com/cloud-ca/cca/pkg/output"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	logutil "sigs.k8s.io/kind/pkg/log"
)

// NewCommand returns a new cobra.Command implementing the root command for cca
func NewCommand() *cobra.Command {
	flags := &cmdutil.GlobalFlags{}
	cmd := &cobra.Command{
		Args:         cobra.NoArgs,
		Use:          "cca",
		Short:        "cca manages authentication, configurations and interactions with the cloud.ca APIs.",
		Long:         "cca manages authentication, configurations and interactions with the cloud.ca APIs.",
		SilenceUsage: true,
		Version:      version.Version(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := flags.Normalize(cmd, viper.Get, args); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&flags.APIURL, "api-url", cmdutil.DefaultAPIURL, "API url cloud.ca resources")
	cmd.PersistentFlags().StringVar(&flags.APIKey, "api-key", "", "API Key to access cloud.ca resources")
	cmd.PersistentFlags().StringVar(&flags.OutputFormat, "output-format", cmdutil.DefaultOutputFormat, "output format "+output.FormatStrings())
	cmd.PersistentFlags().BoolVar(&flags.OutputColor, "output-color", false, "Enable or disable colored output")
	cmd.PersistentFlags().StringVar(&flags.LogLevel, "loglevel", cmdutil.DefaultLogLevel.String(), "log level "+logutil.LevelsString())

	cmd.AddCommand(completion.NewCommand(flags))
	cmd.AddCommand(version.NewCommand(flags))

	return cmd
}

// Run runs the `cca` root command
func Run() error {
	return NewCommand().Execute()
}

// Main wraps Run and sets the log formatter
func Main() {
	// let's explicitly set stdout
	logrus.SetOutput(os.Stdout)

	// this formatter is the default, but the timestamps output aren't
	// particularly useful, they're relative to the command start
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		// we force colors because this only forces over the isTerminal check
		// and this will not be accurately checkable later on when we wrap
		// the logger output with our logutil.StatusFriendlyWriter
		ForceColors: logutil.IsTerminal(logrus.StandardLogger().Out),
	})

	if err := Run(); err != nil {
		os.Exit(1)
	}
}
