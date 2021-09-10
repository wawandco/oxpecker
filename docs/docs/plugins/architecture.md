---
title: "Architecture"
sidebar_position: 1 
---

You might have read a couple of times the `plugin system` that Ox uses, in short, Ox uses a plugin system that allows to add and remove components depending on the need.

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

