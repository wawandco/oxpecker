// tools package contains cmd parts that are specific to tools
// like webpack, packr and other technologies used in Buffalo.
//
// On the long term this package should end up being external,
// once we provide the ability to plug your own tools each of these
// tools should be responsbible to maintain their part.
package tools

import (
	"strings"

	"github.com/gobuffalo/here"
)

// BuildName extracts the last part of the module by splitting on `/`
// this last part is useful for name of the binary and other things.
func BuildName() (string, error) {
	info, err := here.Current()
	if err != nil {
		return "", err
	}

	parts := strings.Split(info.Module.Path, "/")
	name := parts[len(parts)-1]

	return name, nil
}
