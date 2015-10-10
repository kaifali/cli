# Command-line tool for [kaifa.li](http://kaifa.li)

[![Code Climate](https://codeclimate.com/github/kaifali/cli/badges/gpa.svg)](https://codeclimate.com/github/kaifali/cli)

## Install

```shell
$ brew install kaifali/formulae/kaifa
```

## Usage

```shell
usage: kaifa [<flags>] <keyword>

Flags:
  --help           Show context-sensitive help (also try --help-long and --help-man).
  --format="html"  Format: html,json
  --version        Show application version.

Args:
  <keyword>  Keyword of service
```

### Example:

Open browser:

```shell
$ kaifa qiniu
```

Fetch JSON:

```shell
$ kaifa qiniu --format=json
#=> {"keyword":"qiniu","name":"七牛","url":"http://developer.qiniu.com"}
```

## Development

* Install [gpm](https://github.com/pote/gpm)
* `$ gpm install`
* `$ go build kaifa.go`
