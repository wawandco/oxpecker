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
