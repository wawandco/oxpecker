# Oxpecker

Oxpecker is a CLI system we use at Wawandco to wire the plugins we use in our day-to-day development tasks, the functionalities are built into plugins we store in a different repository, this tools allows us to wire those plugins depending on the needs of the project.

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

