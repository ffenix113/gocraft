package types

import (
	"io"

	"github.com/google/uuid"
)

var _ Typer = &UUID{}

type UUID struct {
	Value uuid.UUID
}

func (u *UUID) Read(r io.Reader) {
	io.ReadFull(r, u.Value[:])
}

func (u *UUID) Write(w io.Writer) {
	w.Write(u.Value[:])
}

func (u *UUID) AsString() *String {
	var s String
	s.Value = u.Value.String()
	return &s
}

func NewUUID(data []byte) *UUID {
	var u UUID
	u.Value = uuid.NewMD5(uuid.NameSpaceDNS, data)
	return &u
}

func NewUUIDFromUsername(username string, offline bool) *UUID {
	var u UUID
	if offline {
		username = "OfflinePlayer:" + username
	}
	u.Value = uuid.NewMD5(uuid.NameSpaceDNS, []byte(username))
	return &u
}
