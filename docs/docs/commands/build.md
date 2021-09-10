---
title: "Build"
date: 2021-09-06T14:47:48-05:00
draft: false
type: command
short: "the build command is intended to build the final Go binary for the Ox app, it invokes things like node build process before packing the binary embedding the asset files."
---

Building your app with `ox` is based on the `build` command. You can get more info by running `ox help build`.
### Multiple binaries
One important thing to mention here is that Ox recommended approach is that an application will have multiple binaries built, one of each purpose. That way when a binary is invoked in production it will know the single task it will run.

On a typical app we could have:
```
yourapp
  app
  cmd
    yourapp
      main.go  // The binary that serves the app handlers and routes
    ox
      main.go  // a binary for CLI cron tasks, migrations etc
    worker
      main.go  // a worker binary for things like Temporal or Faktory.
```

And the Dockerfile could just build those to be ready in the Dockerfile.
