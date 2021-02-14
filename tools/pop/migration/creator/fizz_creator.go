package creator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

// FizzCreator model struct for fizz generation files
type FizzCreator struct{}

// Name is the name of the migration type
func (f FizzCreator) Name() string {
	return "fizz"
}

// Create will create 2 .fizz files for the migration
func (f *FizzCreator) Create(dir string, args []string) error {
	name := flect.Underscore(flect.Pluralize(strings.ToLower(args[0])))
	table := fizz.NewTable(name, map[string]interface{}{
		"timestamps": false,
	})

	for _, arg := range args[1:] {
		slice := strings.Split(arg, ":")
		if len(slice) == 1 {
			slice = append(slice, "string")
		}

		o := fizz.Options{}
		name := flect.Underscore(slice[0])
		colType := f.colType(slice[1])

		if name == "id" {
			o["primary"] = true
		}

		if strings.HasPrefix(strings.ToLower(slice[1]), "nulls.") {
			o["null"] = true
		}

		if err := table.Column(name, colType, o); err != nil {
			return err
		}
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, name)

	if err := f.createUPFile(dir, fileName, table); err != nil {
		return err
	}

	if err := f.createDownFile(dir, fileName, table); err != nil {
		return err
	}

	return nil
}

func (f *FizzCreator) createDownFile(dir, name string, table fizz.Table) error {
	filename := fmt.Sprintf("%s.down.fizz", name)
	path := filepath.Join(dir, filename)

	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(table.UnFizz())
	if err != nil {
		return errors.Wrap(err, "error writing file")
	}

	writer.Flush()

	return nil
}

func (f *FizzCreator) createUPFile(dir, name string, table fizz.Table) error {
	filename := fmt.Sprintf("%s.up.fizz", name)
	path := filepath.Join(dir, filename)

	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(table.Fizz())
	if err != nil {
		return errors.Wrap(err, "error writing file")
	}

	writer.Flush()

	return nil
}

func (f *FizzCreator) colType(s string) string {
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
			return f.colType(strings.Replace(s, "nulls.", "", -1))
		}
		return strings.ToLower(s)
	}
}
