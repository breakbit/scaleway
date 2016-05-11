# scaleway

[![Build Status](https://travis-ci.org/breakbit/scaleway.svg?branch=develop)](https://travis-ci.org/breakbit/scaleway) [![Coverage Status](https://coveralls.io/repos/github/breakbit/scaleway/badge.svg?branch=develop)](https://coveralls.io/github/breakbit/scaleway?branch=develop) [![Go Report Card](https://goreportcard.com/badge/github.com/breakbit/scaleway)](https://goreportcard.com/report/github.com/breakbit/scaleway)

scaleway is a Go client library for accessing the [Scaleway API].

This library is written only with the Go standard library, no package dependencies.

## Install

```
go get github.com/breakbit/scaleway
```

## Usage
```go
import "github.com/breakbit/scaleway"
```

Construct a new Scaleway client, then use the various services on the client to
access differents parts of the Scaleway API. For example:


```go
// Create a client
client := scaleway.NewClient(nil)

// Create credentials structure
credentials := scaleway.NewCredentials("foo@bar.com", "foobar")

// Create new token
token, _, _ := client.Tokens.Create(credentials, true)

// Use this token
client.AuthToken = token.ID
```

[Scaleway API]: https://developer.scaleway.com
