package filetypes

import (
	"cuelang.org/go/cue/build"
	"cuelang.org/go/internal/filetypes"
)

type Mode = filetypes.Mode

const (
	Input Mode = iota // The default
	Export
	Def
	Eval
	NumModes
)

func ParseFile(s string, mode Mode) (*build.File, error) {
	return filetypes.ParseFile(s, mode)
}
