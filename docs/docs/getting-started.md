---
sidebar_position: 1
name: getting-started
title: Getting Started
---

In order to get started with the Ox CLI you should install. You can grab the binary from the github repository or you can install from source, which is the recommended way. To install from source you should run:

```sh
go install github.com/wawandco/oxpecker/cmd/ox@latest
```

Once this completes you should have the ox binary in your terminal. You can test it and see if it works by running:

```sh
ox help
```

You should see something like:

```sh
[info] Using wawandco/oxpecker/cmd/ox 

Oxpecker allows to build apps with ease

Usage:
  ox [command]

Commands:
Command      Alias
  help          h       prints help text for the commands registered
  build         b       builds a buffalo app from within the root folder of the project
  dev           d       calls NPM or yarn to start webpack watching the assetst
  db                    database operation commands
  test                  provides the structure for test commands to run and be organized
  fix                   adapts the source code to comply with newer versions of the CLI
  generate      g       Allows to invoke registered generator plugins
  new                   Generates a new app with registered plugins
  task                  Runs grifts tasks previously imported in the CLI
  version       v       returns the current version of Oxpecker CLI
```

Which means you're all set.
