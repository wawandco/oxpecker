package fizz

import "strings"

func columnType(s string) string {
	switch strings.ToLower(s) {
	case "int":
		return "integer"
	case "time.time", "time", "datetime":
		return "timestamp"
	case "uuid.uuid", "uuid":
		return "uuid"
	case "nulls.float32", "nulls.float64":
		return "float"
	case "slices.string", "slices.uuid", "[]string":
		return "varchar[]"
	case "slices.float", "[]float", "[]float32", "[]float64":
		return "numeric[]"
	case "slices.int":
		return "int[]"
	case "slices.map":
		return "jsonb"
	case "float32", "float64", "float":
		return "decimal"
	case "blob", "[]byte":
		return "blob"
	default:
		if strings.HasPrefix(s, "nulls.") {
			return columnType(strings.Replace(s, "nulls.", "", -1))
		}
		return strings.ToLower(s)
	}
}
