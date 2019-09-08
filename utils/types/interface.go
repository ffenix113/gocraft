package types

import "io"

type Typer interface {
	Read(r io.Reader)
	Write(w io.Writer)
}
