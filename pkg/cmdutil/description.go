package cmdutil

import (
	"fmt"
	"strings"

	"github.com/lithammer/dedent"
)

// LongDescription formats long multi-line description and removes
// the left empty space from the lines
func LongDescription(a interface{}) string {
	return strings.TrimLeft(dedent.Dedent(fmt.Sprint(a)), "\n")
}
