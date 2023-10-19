# dashotv/tvdb

TVDB API in Go

[![GoDoc](https://godoc.org/github.com/dashotv/tvdb?status.svg)](https://godoc.org/github.com/dashotv/tvdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/dashotv/tvdb)](https://goreportcard.com/report/github.com/dashotv/tvdb)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## Development

Crate a local `.env` file with the following content:

```
# .env
TVDB_API_URL=https://api4.thetvdb.com/v4
TVDB_API_KEY=your_api_key
TVDB_API_TOKEN=your_api_token
```

Run the following to get the make targets:

> make help

```

Usage:
  make <target>

Targets:
    clean               Remove build related file
  Test:
    test                Run the tests of the project
    coverage            Run the tests of the project and export the coverage
  Lint:
    lint                Run all available linters
    lint-go             Use golintci-lint on your project
  Help:
    help                Show this help.

```
