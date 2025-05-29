package lsp

import (
	"context"

	goplscmd "cuelang.org/go/internal/golangorgx/gopls/cmd"
	"cuelang.org/go/internal/golangorgx/gopls/hooks"
	"cuelang.org/go/internal/golangorgx/tools/tool"
	"cuelang.org/go/internal/lsp/cache"
)

func SetDefaultRegistryFunc(f cache.DefaultRegistryBuildFunc) {
	cache.SetDefaultRegistryFunc(f)
}

func Run(ctx context.Context, args []string) {
	tool.Main(ctx, goplscmd.New(hooks.Options), args)
}
