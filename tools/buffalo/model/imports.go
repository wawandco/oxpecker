package model

import (
	"sort"
	"strings"
)

func buildImports(attrs []attr) []string {
	imps := map[string]bool{"fmt": true}

	ats := attrs
	for _, a := range ats {
		switch a.GoType() {
		case "uuid", "uuid.UUID":
			imps["github.com/gofrs/uuid"] = true
		case "time.Time":
			imps["time"] = true
		default:
			if strings.HasPrefix(a.GoType(), "nulls") {
				imps["github.com/gobuffalo/nulls"] = true
			}
			if strings.HasPrefix(a.GoType(), "slices") {
				imps["github.com/gobuffalo/pop/v5/slices"] = true
			}
		}
	}

	i := make([]string, 0, len(imps))
	for k := range imps {
		i = append(i, k)
	}

	sort.Strings(i)

	return i
}
