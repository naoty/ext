# ext

An interface for command extensions

## Installation

```
$ go get github.com/naoty/ext
```

## Usage

```
$ alias git="ext git"
$ git pr # Run `git-pr` before `git pr`
```

```
$ alias gem="ext gem"
$ gem uninstall all # Run `gem-uninstall-all` before `gem uninstall all`
```

If you want to extend a command by a new subcommand, you should add a command named `<command>-<subcommand>` such as `git-pr`.

## Author

[naoty](https://github.com/naoty)

