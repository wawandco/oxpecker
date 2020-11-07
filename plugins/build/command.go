package build

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/plugins"
)

var _ plugins.Command = (*Command)(nil)

type Command struct {
	builders       []Builder
	afterBuilders  []AfterBuilder
	beforeBuilders []BeforeBuilder
}

func (b Command) Name() string {
	return "build"
}

// Run builds a buffalo app from within the root folder of the project
// To do so, It:
// - Runs NPM or YARN depending on what if finds
// - Runs Packr, Pkger or Other Packing tool
// - Injects database.yml and inflections.
// - Overrides main.go to add migrate
// - Runs go build
func (b *Command) Run(ctx context.Context, root string, args []string) error {

	var err error

	fmt.Println("Before Build:")
	for _, builder := range b.beforeBuilders {
		fmt.Printf(">>> %v BeforeBuilder Running \n\n", builder.Name())

		err = builder.BeforeBuild(ctx, root, args)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err.Error())
			break
		}
	}

	if err == nil {
		fmt.Println("Build:")

		for _, builder := range b.builders {
			fmt.Printf(">>> %v Builder Running \n\n", builder.Name())

			err = builder.Build(ctx, root, args)
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err.Error())
				break
			}
		}
	}

	fmt.Println("After Build:")
	for _, afterBuilder := range b.afterBuilders {
		fmt.Printf(">>> %v AfterBuilder Running \n\n", afterBuilder.Name())

		err = afterBuilder.AfterBuild(root, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {

		if ptool, ok := plugin.(BeforeBuilder); ok {
			b.beforeBuilders = append(b.beforeBuilders, ptool)
		}

		if ptool, ok := plugin.(Builder); ok {
			b.builders = append(b.builders, ptool)
		}

		if ptool, ok := plugin.(AfterBuilder); ok {
			b.afterBuilders = append(b.afterBuilders, ptool)
		}
	}
}
