package dev

import (
	"context"
	"fmt"
	"sync"
)

// dev is the dev command.
type dev struct {
	developers []developer
}

func (d dev) Name() string {
	return "dev"
}

// Run calls NPM or yarn to start webpack watching the assets
// Also starts refresh listening for the changes in Go files.
func (d dev) Run(ctx context.Context, root string, args []string) error {

	var wg sync.WaitGroup
	for _, tool := range d.developers {
		wg.Add(1)
		// TODO: This needs to be parallel.
		go func(t developer) {
			err := t.Develop(ctx, root)
			if err != nil {
				fmt.Println(err)
			}

			wg.Done()
		}(tool)
	}

	wg.Wait()
	return nil
}

func New(tools []interface{}) dev {
	command := dev{}

	for _, tool := range tools {
		if ptool, ok := tool.(developer); ok {
			command.developers = append(command.developers, ptool)
		}
	}

	return command
}
