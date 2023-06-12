# ObjectID

Library to generate MongoDB objectID's

[![Go Report Card](https://goreportcard.com/badge/github.com/Ajnasz/objectid?style=flat-square)](https://goreportcard.com/report/github.com/Ajnasz/objectid)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/Ajnasz/objectid)](https://pkg.go.dev/mod/github.com/Ajnasz/objectid)

## Installation

```sh
go get -u github.com/Ajnasz/objectid
```

### Usage


```go
import github.com/Ajnasz/objectid


func main() {
    oid := objectid.New()
    fmt.Printf("%s", oid)
}
```

## CLI

### Install the cli tool

```sh
go install github.com/Ajnasz/objectid/cmd/objectid
```

or

```sh
make build
bin/objectid
```
