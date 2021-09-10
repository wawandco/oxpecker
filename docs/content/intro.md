---
sidebar_position: 1
name: intro
title: Introduction
---

### Getting Started

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


### Ox vs Buffalo CLI
As mentioned earlier, Ox is not the official CLI for the Go Buffalo development ecosystem, Buffalo provides the `buffalo` command.

Ox is based on the experience we have had at [Wawandco](https://wawand.co) developing sustainable and scalable systems for our clients with Buffalo, where we've evidenced that Buffalo (the library) serves as a huge productivity booster.

We decided to build our own CLI because we don't want to impact others productivity with the choices we've made but we think this could be useful for apps that are starting.

Ox is based on the plugin system that Mark Bates has intended to use in `buffalo-cli`, and allows to add extra plugins based on specific development workflows.

Ox also considers building multiple binaries instead of packing everything in the same binary (how the `buffalo` cli works). See more on the #building section.

### Credits & Acknowledgements
Oxpecker would not be possible without the continuous feedback from the engineering team at [Wawandco](https://wawand.co), the continuous conversations we have inside the company allow us to be always looking for better ways to do things on the CLI.

Also and not less important, this CLI would not be possible without the design that [@markbates](github.com/markbates) did on the `buffalo-cli`, thanks Mark for Buffalo and the Buffalo-cli plugin system.