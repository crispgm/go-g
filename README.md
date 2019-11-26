# go-g

![travis](https://travis-ci.org/crispgm/go-g.svg?branch=master)
![github-workflow-ci](https://github.com/crispgm/go-g/workflows/workflow-ci/badge.svg)
![codecov](https://codecov.io/gh/crispgm/go-g/branch/master/graph/badge.svg)

Go utility packages.

It is named after `g`, so we can call with `g.Foobar` once imported.

## Installation

```shell
go get -u github.com/crispgm/go-g
```

## Usage

### env

* `GetRuntimeEnv`
* `IsDevelopment`
* `IsStaging`
* `IsProduction`

### math

* `Deg2Rad`
* `Rad2Deg`

### net/mac

* `GetMacAddr`

### net/ip

* `GetLocalIP`

### string

* `ExistsInSlice`
* `DedupSlice`

### string/sanitizer

* `RemoveDelimiters`

### string/tokenizer

* `Tokenizer`

## License

MIT
