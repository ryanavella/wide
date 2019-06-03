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
