package modload

import (
	"context"
	"io/fs"

	"cuelang.org/go/internal/mod/modload"
	"cuelang.org/go/mod/modconfig"
	"cuelang.org/go/mod/modfile"
)

func Tidy(ctx context.Context, fsys fs.FS, modRoot string, reg modconfig.Registry) (*modfile.File, error) {
	return modload.Tidy(ctx, fsys, modRoot, reg)
}
