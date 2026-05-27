package s3afero

import (
	"github.com/spf13/afero"
)

type MultiOption func(b *MultiBucketBackend) error

func MultiWithMetaFs(fs afero.Fs) MultiOption { _ = "STUB: not implemented"; return *new(MultiOption) }

func MultiFsFlags(flags FsFlags) MultiOption { _ = "STUB: not implemented"; return *new(MultiOption) }

type SingleOption func(b *SingleBucketBackend) error
