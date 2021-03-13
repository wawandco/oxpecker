package model

var modelTemplate string = `package models

import (
	{{- range $i := .Imports }}
	"{{$i}}"
	{{- end }}
)

// {{ .Name.Proper.String }} model struct
type {{ .Name.Proper.String }} struct {
	{{- range $attr := .Attrs }}
	{{ $attr.Name.Pascalize }}	{{$attr.GoType }} ` + "`" + `json:"{{ $attr.Name.Underscore }}" db:"{{ $attr.Name.Underscore }}"` + "`" + `
	{{- end }}
}

// {{ .Name.Proper.Pluralize }} array model struct of {{ .Name.Proper.String }}
type {{ .Name.Proper.Pluralize }} []{{ .Name.Proper.String }}

// String converts the struct into a string value
func ({{ .Char }} {{ .Name.Proper.String }}) String() string {
	return fmt.Sprintf("%+v\n", {{ .Char }})
}
`

var modelTestTemplate string = `package models

func (ms *ModelSuite) Test_{{ .Name.Proper.String }}() {
	ms.Fail("This test needs to be implemented!")
}`

var modelsBaseTemplate string = `package models

import (
	"bytes"
	"log"
	base "{{.}}"
	
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v5"
)

// Loading connections from database.yml in the pop.Connections
// variable for later usage.
func init() {
	bf, err := base.Config.Find("database.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = pop.LoadFrom(bytes.NewReader(bf))
	if err != nil {
		log.Fatal(err)
	}
}

// DB returns the DB connection for the current environment.
func DB() *pop.Connection {
	c, err := pop.Connect(envy.Get("GO_ENV", "development"))
	if err != nil {
		log.Fatal(err)
	}

	return c
}

`

var modelsTestBaseTemplate string = `package models

import (
	"testing"

	"github.com/gobuffalo/suite/v3"
)

type ModelSuite struct {
	*suite.Model
}

func Test_ModelSuite(t *testing.T) {
	suite.Run(t, &ModelSuite{
		Model: suite.NewModel(),
	})
}

`
