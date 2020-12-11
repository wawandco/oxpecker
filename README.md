# Oxpecker

Oxpecker is a CLI system we use at Wawandco to wire the plugins we use in our day-to-day development tasks, the functionalities are built inside Plugins. 
We store our regular plugins in the [wawandco/oxpecker-plugins](https://github.com/wawandco/oxpecker-plugins) repo. This tools allows us to wire those plugins depending on the needs of the project we're working on.

## Installation

Assuming you have the Go tooling installed and `GOPATH/bin` in your PATH you can install `ox` by running:

```sh
GO111MODULE=on go install github.com/wawandco/oxpecker/cmd/ox
```

## Usage

After installing Ox defaults to have all the plugins in the wawandco/oxpecker-plugins repository, those are based on how we generate/build things. If you want to use your own plugins or pick and choose from that list you can generate cmd/ox/main.go with 

```
ox generate ox
```

Inside that file you can specify the plugins you want to use. You can take a deeper read at how that works in the [plugins docs](docs/PLUGINS.md).

### Help

The help command ships with Oxpecker and allows to get help for a command or subcommand with it. You can invoke it with:

```
ox help [command]
```

For example, `ox help build` displays info about the build command.

```
$ ox help build      
~~~~ Using wawandco/oxpecker/cmd/ox ~~~

[build]
builds a buffalo app from within the root folder of the project

Usage:
  ox build 

Flags:
  -o, --output string   the path the binary will be generated at
      --static          build a static binary using  --ldflags '-linkmode external -extldflags "-static"' (default true)
      --tags strings    tags to pass the go build command
```
