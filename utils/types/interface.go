package types

import "io"

type Typer interface {
	Read(r io.Reader) error
	Write(w io.Writer) error
}
