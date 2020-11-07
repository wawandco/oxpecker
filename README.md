# X

X is a CLI for the Buffalo Framework. The overall goal is to build a simple and opinionated CLI for the Buffalo framework for the tooling that I use everyday, with the possibility to extend to other purposes like plugin system and so on. This source code is heavily inspired by the work @markbates did in the buffalo-cli.

## Installation

To install X you need to run the following command:

```sh
go install github.com/paganotoni/x/cmd/x
```

## Design Constraints

- Simple as possible
- Supports to the tools I use
- Relay in the Go stdlib as much as possible (less dependencies)
- Not limited by other choices made in the Official buffalo CLI.

## Commands

These are the commands that the X CLI contains.

- Dev [DONE]
- Build [IN PROGRESS]
- Test [TODO]
- Help [TODO]
- Fix [TODO][NICE-TO-HAVE]

## Things to cover

- [NEED FOR PRODUCTION] Migrations after built. Need to provide a way to run your migrations.
- BeforeBuild should generate the `:root:/[name].go` file to allow Packr to pack correctly.
- Fixer to move main to `cmd/name`.
- Fixer for models.go and change `models.DB` to `models.DB()` across the app.
- Fixer to translate database.go into `config/database.go` and add `config/constants.go`.
- [NICE TO HAVE] Plugin System.


## Keep in mind

- Only works with Go Modules
- We only build static
- We don't use grifts
- CGO is disabled when building