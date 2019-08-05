package output

import (
	"strings"
)

var outputFormats = []string{"json", "yaml"}

// Get return available output formats
func Get() []string {
	return outputFormats
}

// Has returns true if the list of available output formats
// contains the provided string, false if not
func Has(format string) bool {
	for _, f := range outputFormats {
		if format == f {
			return true
		}
	}
	return false
}

// FormatStrings returns a string representing all output formats.
// this is useful for help text / flag info
func FormatStrings() string {
	var b strings.Builder
	b.WriteString("[")
	for i, format := range outputFormats {
		b.WriteString(format)
		if i+1 != len(outputFormats) {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()

}
