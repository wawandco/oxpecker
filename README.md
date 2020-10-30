# AX

AX is a CLI for the Buffalo framework. The overall goal is to build a simple and opinionated CLI for the Buffalo framework for the tooling that I use everyday, with the possibility to extend to other purposes like plugin system and so on. This source code is heavily inspired by the work @markbates did in the buffalo-cli.

## Design Constraints

- Simple as possible
- Supports to the tools I and My team uses
- Relay in the Go stdlib as much as possible.
- Not limited by other choices made in the Official buffalo CLI.

## Commands

- Build
- Test
- Dev
- Help
- Fix

# Important 

- Only works with Go Modules
- We only build static
- We don't use grifts
- CGO is disabled when building

# Things to cover
- Fixer for models.go and change `models.DB` to `models.DB()` across the app.

