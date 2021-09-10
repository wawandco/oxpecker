package grift

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/markbates/grift/grift"
	"github.com/wawandco/ox/plugins"
)

var (
	// We want grift task to be invoked from any place.
	// this is because of production/compiled usage of the tasks.
	_ plugins.RootFinder = (*Command)(nil)
)

// Grift command is a root command to run tasks
// usage is ox task [name]. If no name is passed this will
// list the tasks
type Command struct{}

func (c Command) Name() string {
	return "task"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "Runs grifts tasks previously imported in the CLI"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		c.list()
		return nil
	}

	task := args[1]
	gc := grift.NewContext(task)
	if len(args) > 2 {
		gc.Args = args[2:]
	}

	return grift.Run(task, gc)
}

func (c Command) list() {
	list := grift.List()
	fmt.Printf("There are %v grift tasks available on this app:\n", len(list))

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 4, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n%s\t%s\t", "task-name", "Full Command")
	fmt.Fprintf(w, "\n%s\t%s\t", "---------", "------------")
	for _, v := range list {
		fmt.Fprintf(w, "\n%v\tox task %s\t", v, v)
	}
	w.Flush()

	fmt.Printf("\n\nrun one of those with: \nox task [task-name]\n")
}

func (c Command) FindRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return wd
}
