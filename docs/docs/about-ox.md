---
sidebar_position: 0
name: about-ox
title: About Ox
---

Ox is a CLI for the Buffalo framework which is focused on making it easy to build and maintain Go applications with a set of  tools that may differ from the official Buffalo CLI. It's based on the [Buffalo](https://github.com/gobuffalo/buffalo) framework, and the [Go](https://golang.org/) programming language. 

One of the big differences between Ox and the official Buffalo CLI is the [plugin System](/docs/plugins/architecture), which allows you to add your own commands and Plugins to the Ox CLI.

### Requirements
Ox requires Go 1.16 or higher.

### Ox and the Buffalo CLI

As mentioned earlier, Ox is not the official CLI for the Go Buffalo development ecosystem, Buffalo provides the `buffalo` command.

Ox is based on the experience we have had at [Wawandco](https://wawand.co) developing sustainable and scalable systems for our clients with Buffalo, where we've evidenced that Buffalo (the library) serves as a huge productivity booster.

We decided to build our own CLI because we don't want to impact others productivity with the choices we've made but we think this could be useful for apps that are starting.

Ox is based on the plugin system that Mark Bates has intended to use in `buffalo-cli`, and allows to add extra plugins based on specific development workflows.

Ox also considers building multiple binaries instead of packing everything in the same binary (how the `buffalo` cli works). See more on the #building section.
