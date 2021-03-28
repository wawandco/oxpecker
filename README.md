# Oxpecker

Oxpecker is an (unofficial) CLI for the Go Buffalo web development ecosystem. Oxpecker provides the `ox` binary which provides commands for common Buffalo development operations.
## Installation

Assuming you have the Go tooling installed and `GOPATH/bin` in your PATH you can install `ox` by running:

```sh
GO111MODULE=on go install github.com/wawandco/oxpecker/cmd/ox
```

## Ox vs Buffalo CLI

As mentioned earlier, Ox is not the official CLI for the Go Buffalo development ecosystem, Buffalo provides the `buffalo` command.

Ox is based on the experience we have had at `Wawandco` developing sustainable and scalable systems for our clients with Buffalo, where we've evidenced that Buffalo (the library) serves as a huge productivity booster.

We decided to build our own CLI because we don't want to impact others productivity with the choices we've made but we think this could be useful for apps that are starting.

Ox is based on the plugin system that Mark Bates has intended to use in `buffalo-cli`, and allows to add extra plugins based on specific development workflows.

Ox also considers building multiple binaries instead of packing everything in the same binary (how the `buffalo` cli works). See more on the #building section.

## Important considerations

Oxpecker bases the way it works on some considerations we've come across in our 3+ years using Buffalo. 

### Requirements
Oxpecker requires:
- [Go Modules]() must be enabled (`GO111MODULE=on`)
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

### New command
### Run command
### DB command
### Generators
### Help
### Plugin System

## Credits & Acknowledgements