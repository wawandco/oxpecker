package resource

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect/name"
	"github.com/pkg/errors"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/tools/buffalo/model"
	"github.com/wawandco/ox/tools/soda/fizz"
)

var (
	//go:embed templates/action.go.tmpl
	actionTemplate string
	//go:embed templates/action_test.go.tmpl
	actionTestTemplate string
	//go:embed templates/index.html.tmpl
	indexHTMLTemplate string
	//go:embed templates/new.html.tmpl
	newHTMLTemplate string
	//go:embed templates/edit.html.tmpl
	editHTMLTemplate string
	//go:embed templates/show.html.tmpl
	showHTMLTemplate string
	//go:embed templates/form.html.tmpl
	formHTMLTemplate string
)

// Resource model struct
type Resource struct {
	Actions  []name.Ident
	Name     name.Ident
	Model    model.Model
	ModelPkg string
	Args     []string

	originalArgs []string
	originalName string
	root         string
}

// New creates a new instance of Resource
func New(root string, args []string) *Resource {
	module := info.ModuleName()
	if module == "" {
		module = root + "/app/models"
	}

	modelsPath := filepath.Join(root, "app", "models")
	model := model.New(modelsPath, args[0], args[1:])
	actions := []name.Ident{
		name.New("list"),
		name.New("show"),
		name.New("new"),
		name.New("create"),
		name.New("edit"),
		name.New("update"),
		name.New("destroy"),
	}

	return &Resource{
		Actions:  actions,
		Args:     args[1:],
		Model:    model,
		ModelPkg: module + "/app/models",
		Name:     name.New(args[0]),

		originalArgs: args[0:],
		originalName: args[0],
		root:         root,
	}
}

// GenerateActions generates the actions for the resource
func (r *Resource) GenerateActions() error {
	actionName := r.Name.Proper().Pluralize().Underscore().String()
	dirPath := filepath.Join(r.root, "app", "actions")
	actions := map[string]string{
		actionName:           actionTemplate,
		actionName + "_test": actionTestTemplate,
	}

	for name, content := range actions {
		filename := name + ".go"
		path := filepath.Join(dirPath, filename)
		err := source.Build(path, content, r)
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateModel generates the migrations for the resource
func (r *Resource) GenerateMigrations() error {
	migrationPath := filepath.Join(r.root, "migrations")
	creator := fizz.Creator{}

	if err := creator.Create(migrationPath, r.originalName, r.originalArgs); err != nil {
		return errors.Wrap(err, "failed creating migrations")
	}

	return nil
}

// GenerateModel generates the model for the resource
func (r *Resource) GenerateModel() error {
	if err := r.Model.Create(); err != nil {
		return errors.Wrap(err, "error creating model")
	}

	return nil
}

// GenerateModel generates the templates for the resource
func (r *Resource) GenerateTemplates() error {
	templates := map[string]string{
		"index": indexHTMLTemplate,
		"new":   newHTMLTemplate,
		"edit":  editHTMLTemplate,
		"show":  showHTMLTemplate,
		"form":  formHTMLTemplate,
	}

	dirPath := filepath.Join(r.root, "app", "templates", r.Name.Underscore().String())
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	for name, content := range templates {
		filename := name + ".plush.html"
		path := filepath.Join(dirPath, filename)

		err := source.Build(path, content, r)
		if err != nil {
			return err
		}
	}

	return nil
}
