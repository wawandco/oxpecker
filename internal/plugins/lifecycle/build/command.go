package build

import (
	"context"
	"fmt"

	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)

type Command struct {
	buildPlugins []plugins.Plugin

	builders       []Builder
	afterBuilders  []AfterBuilder
	beforeBuilders []BeforeBuilder
}

func (b Command) Name() string {
	return "build"
}

//HelpText resturns the help Text of build function
func (b Command) HelpText() string {
	return "builds a buffalo app from within the root folder of the project"
}

// Run builds a buffalo app from within the root folder of the project
// To do so, It:x
// - Runs NPM or YARN depending on what if finds
// - Runs Packr, Pkger or Other Packing tool
// - Injects database.yml and inflections.
// - Overrides main.go to add migrate
// - Runs go build
func (b *Command) Run(ctx context.Context, root string, args []string) error {
	var err error
	for _, builder := range b.beforeBuilders {
		fmt.Printf(">>> %v BeforeBuilder Running \n", builder.Name())

		err = builder.RunBeforeBuild(ctx, root, args)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err.Error())
			break
		}
	}

	if err == nil {
		for _, builder := range b.builders {
			fmt.Printf(">>> %v Builder Running \n", builder.Name())

			err = builder.Build(ctx, root, args)
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err.Error())
				break
			}
		}
	}

	for _, afterBuilder := range b.afterBuilders {
		fmt.Printf(">>> %v AfterBuilder Running \n", afterBuilder.Name())

		err = afterBuilder.RunAfterBuild(root, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		isBuildPlugin := false
		if ptool, ok := plugin.(BeforeBuilder); ok {
			isBuildPlugin = true
			b.beforeBuilders = append(b.beforeBuilders, ptool)
		}

		if ptool, ok := plugin.(Builder); ok {
			isBuildPlugin = true
			b.builders = append(b.builders, ptool)
		}

		if ptool, ok := plugin.(AfterBuilder); ok {
			isBuildPlugin = true
			b.afterBuilders = append(b.afterBuilders, ptool)
		}

		if isBuildPlugin {
			b.buildPlugins = append(b.buildPlugins, plugin)
		}
	}
}
