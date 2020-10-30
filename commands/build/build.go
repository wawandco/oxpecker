package build

import (
	"context"
	"fmt"
)

type build struct {
	builders       []builder
	afterBuilders  []afterBuilder
	beforeBuilders []beforeBuilder
}

func (b build) Name() string {
	return "build"
}

// Run builds a buffalo app from within the root folder of the project
// To do so, It:
// - Runs NPM or YARN depending on what if finds
// - Runs Packr, Pkger or Other Packing tool
// - Injects database.yml and inflections.
// - Overrides main.go to add migrate
// - Runs go build
func (b build) Run(ctx context.Context, root string, args []string) error {

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

func New(tools []interface{}) build {
	command := build{}

	for _, tool := range tools {

		if ptool, ok := tool.(beforeBuilder); ok {
			command.beforeBuilders = append(command.beforeBuilders, ptool)
		}

		if ptool, ok := tool.(builder); ok {
			command.builders = append(command.builders, ptool)
		}

		if ptool, ok := tool.(afterBuilder); ok {
			command.afterBuilders = append(command.afterBuilders, ptool)
		}
	}

	return command
}
