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

// Package cmdutil contains general utility of the cca command
package cmdutil

import (
	"github.com/cloud-ca/cca/pkg/output"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GlobalFlags for the cca command
type GlobalFlags struct {
	APIURL        string
	APIKey        string
	ColorOutput   bool
	EnvironmentID string
	LogLevel      string
	OutputColor   bool
	OutputFormat  string
}

// Normalize checks and normalizes input flags and falls back to default values when needed
func (gf *GlobalFlags) Normalize(cmd *cobra.Command, fn func(key string) interface{}, args []string) error {
	if err := gf.parseLogLevel(cmd, args); err != nil {
		return err
	}
	if err := gf.parseColorOutput(cmd, args); err != nil {
		return err
	}
	if err := gf.parseOutputFormat(cmd, args); err != nil {
		return err
	}

	return nil
}

func (gf *GlobalFlags) parseLogLevel(cmd *cobra.Command, args []string) error {
	level := DefaultLogLevel
	parsed, err := logrus.ParseLevel(gf.LogLevel)
	if err != nil {
		logrus.Warnf("Invalid log level '%s', defaulting to '%s'", gf.LogLevel, level)
	} else {
		level = parsed
	}
	logrus.SetLevel(level)
	return nil
}

func (gf *GlobalFlags) parseColorOutput(cmd *cobra.Command, args []string) error {
	nocolor, err := cmd.Flags().GetBool("nocolor")
	if err != nil {
		logrus.Warnf("Invalid nocolor value '%v', defaulting to '%v'", nocolor, false)
	} else {
		gf.ColorOutput = !nocolor
	}
	return nil
}

func (gf *GlobalFlags) parseOutputFormat(cmd *cobra.Command, args []string) error {
	if !output.Has(gf.OutputFormat) {
		logrus.Warnf("Invalid output format '%s', defaulting to '%s'", gf.OutputFormat, DefaultOutputFormat)
		gf.OutputFormat = DefaultOutputFormat
	}
	return nil
}
