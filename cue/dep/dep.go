package dep

import (
	"context"
	"iter"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/internal/core/dep"
	"cuelang.org/go/internal/core/eval"
	"cuelang.org/go/internal/value"
)

func Deps(v cue.Value) iter.Seq2[cue.Path, error] {
	r, n := value.ToInternal(v)
	ctx := eval.NewContext(r, n)

	return func(yield func(cue.Path, error) bool) {
		if err := dep.Visit(&dep.Config{Dynamic: true}, ctx, n, func(d dep.Dependency) error {
			depValue := value.Make(ctx, d.Node)

			if d.Import() != nil {
				return nil
			}

			if !d.IsRoot() {
				return nil
			}

			if !yield(depValue.Path(), nil) {
				return context.Canceled
			}

			return nil
		}); err != nil {
			if !errors.Is(err, context.Canceled) {
				yield(cue.MakePath(), err)
			}
		}
	}
}
