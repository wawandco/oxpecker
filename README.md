# Oxpecker

Oxpecker is a CLI for web applications being build with Go and the Buffalo Framework. It aims to provide support for the everyday development using Go and the Buffalo stack.

## Installation

Assuming you have the Go tooling installed and `GOPATH/bin` in your PATH you can install `ox` by running:

```sh
GO111MODULE=on go install github.com/paganotoni/oxpecker/cmd/ox
```

Or using gobinaries with:

```sh
curl -sf https://gobinaries.com/paganotoni/oxpecker/cmd/ox | sh
```

## Usage

```
ox [command]

p.e 
ox dev
ox build
ox test
```

## Commands

Commands are loaded from plugins, instead of being a hardcoded list of commands plugins used in the CLI will provide the commands that will be available at the CLI runtime. CLI will identify those commands with the Command interface.

## Plugins

TODO: Explain how this will work.

## Why another CLI?

TLDR: I want to. And in doing so want to avoid discussions about previous choices made in the v1 Buffalo CLI.

**Long version**

To explain the Why of this CLI tool I have to mention that there is a new CLI for buffalo that's being developed, I personally have been working on it.

While Working on Buffalo-cli I learnt a lot of things that are done and took a lot of ideas for the design/implementation of this CLI tool. I also noticed that there are some patterns that repeat in the Go/Buffalo code-bases I work that would like to incorporate into the CLI, but I could not do that with the Buffalo-cli because I don't have the freedom to break everybody's code there.

So I thought about building this CLI as a way to try out those ideas I and my team have had by incorporating some of the lesson learned on both the Buffalo-cli [LINK] and the current Buffalo cli.

## Principles

- Guided by experience (Extracting is preferred over Imagining).
- Prefer Go: 
    - Want to use the Go standard library as much as possible
    - Avoid YML/TOML/Other and other markup languages for configuration.
    - Embrace Go modules and require it
- Keep it simple.
- Convention over configuration.

## Open Topics

- CGO is disabled, should it be? Does this means we are not supporting sqlite ?
- Should we include tasks? (grift/jim/other?)
- Need to do integration testing for the CLI

