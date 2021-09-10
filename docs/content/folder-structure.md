---
title: "Folder Structure"
date: 2021-09-02T14:45:49-05:00
draft: false
sidebar_position: 2
---

Applications generated with Ox provide a folder structure that serves as a base for different operations in the development workflow as well as a layout for the code base, an Ox generated application has the following folder structure:

```sh
yourapp
  app
    actions               # Buffalo handlers in the app
    assets                # The JS/CSS/Images assets for the app
    middleware            # Middleware used in the app
    models                # Application models
    render                # Render engine and helpers
    tasks                 # Tasks
    templates             # Plush templates used in the app
    app.go                # Constructor for the app (app.New method)
    routes.go             # App routes in the setRoutes(app) method
  cmd                     
    yourapp               
      main.go             # This is the main application binary
  config                  
    database.yml          # Database configuration
  migrations              # Database migrations folder
  public                  # Built assets end here
  .babelrc                # Config for babel (js tooling)
  .buffalo.dev.yml        # Config for Refresh
  Dockerfile              # The dockerfile to build the app Docker image
  embed.go                # Embedded files configuration
  go.mod                  # Application module and deps definition
  package.json            # JS dependencies
  postcss.config.js       # PostCSS config
  webpack.config.js       # Webpack bundler configuration
```
