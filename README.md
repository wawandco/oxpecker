# X

X is a CLI for web applications being build with Go and the Buffalo Framework. It aims to provide support for the everyday operations of developers using that stack.

## Installation

To install X you need to run the following command:

```sh
go install github.com/paganotoni/x/cmd/x
```

## Commands

These are the commands that the X CLI contains.

- Dev [DONE]
- Build [IN PROGRESS]
- Test [TODO]
- Help [TODO]
- Fix [TODO][NICE-TO-HAVE]

## Things to cover

- [NEED FOR PRODUCTION] Migrations after built. Need to provide a way to run your migrations.
- Packr needs to check that there is a go file in the root package otherwise it will not work (generate `:root:/[name].go` file to allow Packr to pack correctly).
- Fixer to move main to `cmd/name`.
- Fixer for models.go and change `models.DB` to `models.DB()` across the app.
- Fixer to translate database.go into `config/database.go` and add `config/constants.go`.
- [NICE TO HAVE] Plugin System.


## Why another CLI?

TLDR: Because I can! (And want to have this experience of designing and using a CLI for the stack that I and [my teams](https://wawand.co) work on everyday)

**Long version**

To explain the Why of this CLI tool I have to mention that there is a new CLI for buffalo that's being developed, I personally have been working on it.

While Working on Buffalo-cli I learnt a lot of things that are done and took a lot of ideas for the design/implementation of this CLI tool. I also noticed that there are some patterns that repeat in the Go/Buffalo codebases I work that I would like to incorporate into the CLI, but I could not do that with the Buffalo-cli because I don't have the freedom to break everybody's code there.

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

- Build is currently only doing the `--static` part. Should that be an option ?
- CGO is disabled, should it be?
- Should we include tasks? (grift/jim/other?)

