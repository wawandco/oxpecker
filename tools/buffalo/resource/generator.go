package resource

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wawandco/ox/internal/log"
)

// Generator allows to identify resource as a plugin
type Generator struct{}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "buffalo/generate-resource"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
	return "resource"
}

// Generate generates the actions, model, templates and migration files from the given attrs
// app/actions/[name].go
// app/actions/[name]_test.go
// app/models/[name].go
// app/models/[name]_test.go
// app/templates/[name]/index.plush.html
// app/templates/[name]/new.plush.html
// app/templates/[name]/edit.plush.html
// app/templates/[name]/show.plush.html
// migrations/[name].up.fizz
// migrations/[name].down.fizz
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate resource [name]`")
	}

	resource := New(root, args[2:])

	// Generating Templates
	log.Info("Generating Actions...\n")
	if err := resource.GenerateActions(); err != nil {
		return errors.Wrap(err, "generating actions error")
	}

	// Generating Templates
	log.Info("Generating Templates...\n")
	if err := resource.GenerateTemplates(); err != nil {
		return errors.Wrap(err, "generating templates error")
	}

	// Generating Model
	log.Info("Generating Model...\n")
	if err := resource.GenerateModel(); err != nil {
		return errors.Wrap(err, "generating model error")
	}

	// // Generating Migration
	log.Info("Generating Migrations...\n")
	if err := resource.GenerateMigrations(); err != nil {
		return errors.Wrap(err, "generating migrations error")
	}

	log.Infof("%s resource has been generated successfully \n", resource.originalName)

	return nil
}
