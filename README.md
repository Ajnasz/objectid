# ObjectID

Library to generate MongoDB objectID's

[![Go Report Card](https://goreportcard.com/badge/github.com/Ajnasz/objectid?style=flat-square)](https://goreportcard.com/report/github.com/Ajnasz/objectid)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/Ajnasz/objectid)](https://pkg.go.dev/mod/github.com/Ajnasz/objectid)

## Command usage

```
$ objectid -h
Usage of objectid:
  -format string
        format of objectid: hex, base64 (default "hex")
  -from-time string
        create a new objectid from a date time (RFC3339, $(date -I), $(date -Ihours), $(date -d -Iminutes), $(date -Iseconds)
  -n int
        number of objectid to generate (default 1)
  -separator string
        separator between objectids (default "\n")
  -to-time string
        convert objectid to time
```

### Generate object id

```sh
$ objectid
6435f4800000000000000000
```

### Generate object id in base64

```sh
$ objectid -format base64
ZiesvNpsbBz8QUGk
```

### Generate multiple object id

```sh
$ objectid -n 3
6435f4800000000000000000
6435f4800000000000000001
6435f4800000000000000002
```


### Object id from time

```sh
$ objectid -from-time 2023-04-12
6435f4800000000000000000%
```

### Object id to time

```sh
$ objectid -to-time 6435f4800000000000000000
2023-04-12 00:00:00 +0000 UTC
```

## Library Usage


```sh
go get -u github.com/Ajnasz/objectid
```

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
