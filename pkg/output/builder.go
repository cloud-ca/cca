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

// Package output contains helper utility to build and format the output
package output

// Builder is used to prepare the output. It internally
// create a Formatter and use it to print the 'object'
// to different formats and colors (based on the flags)
type Builder struct {
	format  string
	colored bool
}

// NewBuilder returns a new output.Builder with desired format and colored output
func NewBuilder(format string, colored bool) *Builder {
	return &Builder{
		format:  format,
		colored: colored,
	}
}

// Build builds the callback function to be used directly in cobra.Command
// in order not to pass around private structs from go-cloudca library
func (b *Builder) Build(fn func(*Formatter) error) error {
	formatter := &Formatter{
		builder: b,
	}
	return fn(formatter)
}