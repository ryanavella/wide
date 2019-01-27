# Wide [![](https://godoc.org/github.com/ryanavella/wide?status.svg)](https://godoc.org/github.com/ryanavella/wide)

Uint128 and Int128 for Go.

Licensed under the Unlicense (https://unlicense.org/)

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
