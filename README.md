# Oxpecker

Oxpecker is an (unofficial) CLI for the Go Buffalo web development ecosystem. Oxpecker provides the `ox` binary which provides commands for common Buffalo development operations.
## Important considerations

Oxpecker bases the way it works on some considerations we've come across in our 3+ years using Buffalo. 

### Requirements
Oxpecker requires:
- [Go Modules](https://blog.golang.org/using-go-modules) must be enabled (`GO111MODULE=on`)
- Go 1.16.x or newer. This is because oxpecker generated code is based on the embed package. 
### Folder structure
Oxpecker considers a different folder structure than the `buffalo` CLI. So its important to describe it, a typical oxpecker app looks like the following:

```
yourapp
  app
    actions               // Buffalo handlers in the app
    assets                // The JS/CSS/Images assets for the app
    middleware            // Middleware used in the app
    models                // Application models
    render                // Render engine and helpers
    tasks                 // Grift tasks
    templates             // Plush templates used in the app
    app.go                // Constructor for the app (app.New method)
    routes.go             // App routes in the setRoutes(app) method
  cmd
    yourapp
      main.go             // This is the main application binary
  config
    database.yml          // Database configuration
  migrations              // Database migrations
  public                  // Built assets end here
  .babelrc                // Config for babel (js tooling)
  .buffalo.dev.yml        // Config for Refresh
  Dockerfile               // The dockerfile to build the app Docker image
  embed.go                // Embedded files configuration
  go.mod                  // Application module and deps definition
  package.json            // JS dependencies
  postcss.config.js        // PostCSS config
  webpack.config.js        // Webpack bundler configuration
```
### Building
Building your app with `ox` is based on the `build` command. You can get more info by running `ox help build`.

One important thing to mention here is that Ox considers an application will have multiple binaries built, so it does not try to pack everything on the same binary.

On a typical app we could have:
```
yourapp
  app
  cmd
    yourapp
      main.go  // The binary that serves the app handlers and routes
    cli
      main.go  // a binary for CLI cron tasks, migrations etc
    worker
      main.go  // a worker binary for things like Temporal.io or Faktory.
```

And the Dockerfile could just build those to be ready in the Dockerfile. 
### Help
The help command serves as a live documentation for each of the commands in the Oxpecker CLI. You can see the top level help by running:

```
ox help
```

Also, you can get specific help for a particular command by running `ox help [command]`. For example:

```
ox help new
```
## Getting started
### Installing the CLI

Assuming you have the Go tooling installed and `GOPATH/bin` in your PATH you can install `ox` by running:

```sh
GO111MODULE=on go install github.com/wawandco/oxpecker/cmd/ox
```

### Creating your app

Ox provides the `new` command to generate a new application from the ground, the command receives the name of the app as the first and required argument.

```
ox new yourapp
```

By running this command you will initialize your application codebase.

### Setting up your Database

One important step for getting started is to create your development database, to do so you will need to run:

```
ox db create
```

And Ox will instruct your DBMS to create the database. 
### Running your application

Once the application codebase and the database are build you can start your application by running:

```
ox dev
```

This command will start your application on `http://127.0.0.1:3000` on development mode.

## Plugin System

You might have read a couple of times the `plugin system` that Oxpecker uses, in short, Oxpecker uses a plugin system that allows to add and remove components depending on the need.

### Base plugins

To start Oxpecker uses a [base set of plugins](https://github.com/wawandco/oxpecker/blob/da3802e39c839864827d693f0fa6c2339626b0cb/tools/tools.go#L44), these include the common things used on application development with Buffalo.

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

    "github.com/wawandco/oxpecker/cli"
    "github.com/wawandco/oxpecker/tools/soda"
)

// main function for the tooling cli, will be invoked by Oxpecker
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

### Ox vs Buffalo CLI

As mentioned earlier, Ox is not the official CLI for the Go Buffalo development ecosystem, Buffalo provides the `buffalo` command.

Ox is based on the experience we have had at [Wawandco](https://wawand.co) developing sustainable and scalable systems for our clients with Buffalo, where we've evidenced that Buffalo (the library) serves as a huge productivity booster.

We decided to build our own CLI because we don't want to impact others productivity with the choices we've made but we think this could be useful for apps that are starting.

Ox is based on the plugin system that Mark Bates has intended to use in `buffalo-cli`, and allows to add extra plugins based on specific development workflows.

Ox also considers building multiple binaries instead of packing everything in the same binary (how the `buffalo` cli works). See more on the #building section.
## Credits & Acknowledgements

Oxpecker would not be possible without the continuous feedback from the engineering team at [Wawandco](https://wawand.co), the continuous conversations we have inside the company allow us to be always looking for better ways to do things on the CLI.

Also and not less important, this CLI would not be possible without the design that [@markbates](github.com/markbates) did on the `buffalo-cli`, thanks Mark for Buffalo and the Buffalo-cli plugin system.