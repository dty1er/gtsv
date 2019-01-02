# gtsv - Fast TSV Parser written in Go

[![Build Status](https://travis-ci.org/yagi5/gtsv.svg?branch=master)](https://travis-ci.org/yagi5/gtsv)
[![Coverage Status](https://coveralls.io/repos/github/yagi5/gtsv/badge.svg?branch=master)](https://coveralls.io/github/yagi5/gtsv?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/yagi5/gtsv)](https://goreportcard.com/report/github.com/yagi5/gtsv)
[![GoDoc](https://godoc.org/github.com/yagi5/gtsv?status.svg)](https://godoc.org/github.com/yagi5/gtsv)

### Installation

```shell
$ go get -u github.com/yagi5/gtsv
```

### Features

* get values as specific type
* get row and column numbers which error caused 

### Usage

Valid TSV case (column numbers of every row are all the same, and type specify is compatible) :

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/yagi5/gtsv"
)

func main() {
  tsv := "1\t2.1\ta\n" +
    "4\t5.2\tb\n" +
    "7\t8.3\tc\n"

  gt := gtsv.New(bytes.NewBufferString(tsv)) // pass io.Reader

  for gt.Next() {
    fmt.Println(gt.Int())
    fmt.Println(gt.Float64())
    fmt.Println(gt.String())
  }
  fmt.Println(gt.Error())
}
```

Output is:

```
1
2.1
a
4
5.2
b
7
8.3
c
<nil>
```

Invalid TSV Case:

```go
func main() {
  tsvInvalid := "1\t2.1\ta\n" +
    "4\t5.2\tb\n" +
    "a\t8.3\tc\n" // first column is not int

  gt := gtsv.New(bytes.NewBufferString(tsvInvalid))

  for gt.Next() {
    fmt.Println(gt.Int())
    fmt.Println(gt.Float64())
    fmt.Println(gt.String())
  }

  if err := gt.Error(); err != nil {
    er := err.(gtsv.Error)
    fmt.Printf("error row: %d, col: %d", er.Row(), er.Col())
	}
}
```

Output is:

```
Output: error row: 3, col: 1
```

For more detail, see [godoc](https://godoc.org/github.com/yagi5/gtsv).

### Lisence

MIT
