# dashotv/tvdb

TVDB API in Go

[![Build Status](https://drone.dasho.net/api/badges/dashotv/tvdb/status.svg?ref=refs/heads/main)](https://drone.dasho.net/dashotv/tvdb)
[![GoDoc](https://godoc.org/github.com/dashotv/tvdb?status.svg)](https://godoc.org/github.com/dashotv/tvdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/dashotv/tvdb)](https://goreportcard.com/report/github.com/dashotv/tvdb)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## Usage

Install the package with:

```bash
go get github.com/dashotv/tvdb
```

Import the package with:

```go
import "github.com/dashotv/tvdb"
```

Create a new client with:

```go
client := tvdb.New(apikey)
```

If you don't already have a token, you can obtain one with:

```go
// Authenticate with your API key. This will return the token and
// will also configure the client with Bearer authentication.
token, err := client.Login()
if err != nil {
    // handle error
}
```

You should store the token somewhere, by default the token is viable for 30 days. TVDB doesn't appear to care if you auth every call, but it adds a lot of overhead.

If you already have the token, you can set it with:

```go
client.SetToken(token)
```

This will configure the client with the token and configure Bearer authentication.

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
