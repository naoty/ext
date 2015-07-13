# ext

An interface for command extensions

## Installation

```
$ go get github.com/naoty/ext
```

## Usage

`ext` provides a interface for command extensions. If you want to extend a command by a new subcommand, you should add a command named `<command>-<subcommand>` such as `git-pr`. `ext` will look up and run the extension before an original command.

```
$ alias git="ext git"
$ git pr
```

If `git-pr` is found, `ext` will run `git-pr` instead `git pr`.

```
$ alias gem="ext gem"
$ gem uninstall all
```

`ext` looks up and runs a binary in following order: `gem-uninstall-all`, `gem-uninstall all`, `gem uninstall all`.

## Author

[naoty](https://github.com/naoty)

