[![](https://godoc.org/github.com/ryanavella/wide?status.svg)](https://godoc.org/github.com/ryanavella/wide) [![Go Report Card](https://goreportcard.com/badge/github.com/ryanavella/wide)](https://goreportcard.com/report/github.com/ryanavella/wide) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/ryanavella/wide/blob/develop/LICENSE-MIT) [![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](https://github.com/ryanavella/wide/blob/develop/LICENSE-UNLICENSE)

# Wide

Uint128 and Int128 for Go.

Wide is free and open source software distributed under the terms of both the MIT License and the Unlicense.

## Installing

```shell
go get github.com/ryanavella/wide
```

## Usage

```golang
package main

import (
	"fmt"

	"github.com/ryanavella/wide"
)

func main() {
	a := wide.Int128FromInt64(-3)
	b := wide.Int128FromInt64(2)
	fmt.Println(a, b, a.Add(b), a.Sub(b), a.Mul(b), a.Div(b), a.Mod(b))
}
```

## Scope

This package is intended for efficient and fast computations (i.e. for scientific and mathematical applications). There are no plans to support applications which require constant-time cryptographic security.

## Contributions

See [contributor guidelines](CONTRIBUTING.md).
