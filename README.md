# Command-line tool for [kaifa.li](http://kaifa.li)

## Install

```shell
$ brew install kaifali/formulae/kaifa
```

## Usage

```shell
usage: kaifa [<flags>] <keyword>

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Args:
  <keyword>  Keyword of service
```

Example:

```shell
$ kaifa qiniu
```

## Development

* Install [gpm](https://github.com/pote/gpm)
* `$ gpm install`
* `$ go build kaifa.go`
