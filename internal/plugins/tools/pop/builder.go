package pop

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/paganotoni/x/internal/plugins/tools/pop/templates"
)

func (p *Plugin) RunBeforeBuild(ctx context.Context, root string, args []string) error {
	f, err := p.findConfig()
	if err != nil {
		return err
	}

	content, err := p.buildDatabaseInit(f)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join("config", "gen_database.go"), []byte(content), 0600)
}

func (p *Plugin) findConfig() ([]byte, error) {
	f, err := ioutil.ReadFile(filepath.Join("config", "database.yml"))
	if err != nil {
		return []byte{}, err
	}

	return f, err
}

func (p *Plugin) buildDatabaseInit(fileContent []byte) (string, error) {
	tmpl, err := template.New("config").Parse(templates.DatabaseConfiguration)
	if err != nil {
		return "", err
	}

	bb := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(bb, struct {
		Config string
	}{
		Config: string(fileContent),
	})

	if err != nil {
		return "", err
	}

	dat, err := ioutil.ReadAll(bb)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func (p *Plugin) RunAfterBuild(root string, args []string) error {
	err := os.Remove(filepath.Join("config", "gen_database.go"))
	if os.IsNotExist(err) {
		return nil
	}

	return err
}
