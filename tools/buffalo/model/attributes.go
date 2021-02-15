package model

import (
	"strings"

	"github.com/gobuffalo/flect/name"
)

type attr struct {
	Name       name.Ident
	CommonType string
}

// GoType returns the Go type for an Attr based on its type
func (a attr) GoType() string {
	switch strings.ToLower(a.CommonType) {
	case "text":
		return "string"
	case "timestamp", "datetime", "date", "time":
		return "time.Time"
	case "nulls.bool":
		return "nulls.Bool"
	case "nulls.int":
		return "nulls.Int"
	case "nulls.decimal", "nulls.float":
		return "nulls.Float64"
	case "nulls.text":
		return "nulls.String"
	case "nulls.time":
		return "nulls.Time"
	case "nulls.uuid":
		return "nulls.UUID"
	case "uuid":
		return "uuid.UUID"
	case "json", "jsonb":
		return "slices.Map"
	case "[]string":
		return "slices.String"
	case "[]int":
		return "slices.Int"
	case "slices.float", "[]float", "[]float32", "[]float64":
		return "slices.Float"
	case "decimal", "float":
		return "float64"
	case "[]byte", "blob":
		return "[]byte"
	default:
		return a.CommonType
	}
}

func buildAttrs(args []string) []attr {
	var attrs []attr
	defaults := defaultAttrs(args)
	args = append(defaults, args...)

	for _, arg := range args {
		slice := strings.Split(arg, ":")
		if len(slice) == 1 {
			slice = append(slice, "string")
		}

		attrs = append(attrs, attr{
			Name:       name.New(slice[0]),
			CommonType: strings.ToLower(slice[1]),
		})
	}

	return attrs
}

// defaultAttrs appends the default attributes if they are not specified
func defaultAttrs(args []string) []string {
	defaults := []string{"id:uuid", "created_at:timestamp", "updated_at:timestamp"}

	if len(args) == 0 {
		return defaults
	}

	m := make(map[string]bool)
	for _, arg := range args {
		element := strings.ToLower(strings.Split(arg, ":")[0])
		if _, ok := m[element]; !ok {
			m[element] = true
		}
	}

	attrs := []string{}
	if _, ok := m["id"]; !ok {
		attrs = append(attrs, defaults[0])
	}

	if _, ok := m["created_at"]; !ok {
		attrs = append(attrs, defaults[1])
	}

	if _, ok := m["updated_at"]; !ok {
		attrs = append(attrs, defaults[2])
	}

	return attrs
}
