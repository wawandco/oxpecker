---
title: "Plugins"
date: 2021-09-02T14:47:08-05:00
draft: false
sidebar_position: 9
---

## Plugin System

You might have read a couple of times the `plugin system` that Ox uses, in short, Ox uses a plugin system that allows to add and remove components depending on the need.

### Base plugins

To start Ox uses a [base set of plugins](https://github.com/wawandco/ox/blob/da3802e39c839864827d693f0fa6c2339626b0cb/tools/tools.go#L44), these include the common things used on application development with Buffalo.

- Pop
- Soda
- Flect
- Envy
- Tags
- Validate
- Grift

### Customizing through plugins
Oftentimes you may need to have your own CLI commands for common operations for your team. While the base plugins provide a foundation for Buffalo development these may fall short for specific team choices.

In these cases, ox provides a plugin system that may come handy, to use it you will need to generate `cmd/cli/main.go` by running (within your app folder):

```
ox generate cli
```

This will generate a file on `yourapp/cmd/cli/main.go` that will look something like:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "yourapp"
    _ "yourapp/app/tasks"
    _ "yourapp/app/models"

    "github.com/wawandco/ox/cli"
    "github.com/wawandco/ox/tools/soda"
)

// main function for the tooling cli, will be invoked by Ox
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its lifecycle.
func main() {
    // using Soda Plugins
    cli.Use(soda.Plugins(yourapp.Migrations)...)
    err := cli.Run(context.Background(), os.Args)
    if err != nil {
        fmt.Printf("[error] %v \n", err.Error())

        os.Exit(1)
    }
}
```

As you can see, the CLI instance allows to specify the plugins you want to use, and uses Base plugins to start. In order to use your own plugin you would just have to add those to the plugins that the CLI will use.

### Building your CLI
One important part to mention is that the build command will not attempt to build the CLI folder. Instead the developer will need to do it when needed, either on the CLI

Either on your Dockerfile or your build system you should include something like `go build ./cmd/cli` to ensure that the cli binary gets to your running environment.