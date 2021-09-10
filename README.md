# Ox

### Installation
Ox is an (unofficial) CLI for the Go Buffalo web development ecosystem. The Ox binary which provides commands for common Buffalo development operations.

To install ox, run:

```sh
go install github.com/wawandco/ox/cmd/ox@latest
```

To know more about Ox take a look at [our documentation site](https://oxcli.dev).

### Ox vs Buffalo CLI

As mentioned earlier, Ox is not the official CLI for the Go Buffalo development ecosystem, Buffalo provides the `buffalo` command.

Ox is based on the experience we have had at [Wawandco](https://wawand.co) developing sustainable and scalable systems for our clients with Buffalo, where we've evidenced that Buffalo (the library) serves as a huge productivity booster.

We decided to build our own CLI because we don't want to impact others productivity with the choices we've made but we think this could be useful for apps that are starting.

Ox is based on the plugin system that Mark Bates has intended to use in `buffalo-cli`, and allows to add extra plugins based on specific development workflows.

Ox also considers building multiple binaries instead of packing everything in the same binary (how the `buffalo` cli works). See more on the #building section.