package s3afero

import (
	"io"

	"github.com/johannesboyne/gofakes3"
	"github.com/spf13/afero"
)

var emptyPrefix = &gofakes3.Prefix{}

type readerWithCloser struct {
	io.Reader
	closer func() error
}

var _ io.ReadCloser = &readerWithCloser{}

func limitReadCloser(rdr io.Reader, closer func() error, sz int64) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (rwc *readerWithCloser) Close() error { _ = "STUB: not implemented"; return nil }

// ensureNoOsFs makes a best-effort attempt to ensure you haven't used
// afero.OsFs directly in any of these backends; to do so would risk exposing
// you to RemoveAll against your `/` directory.
func ensureNoOsFs(name string, fs afero.Fs) error { _ = "STUB: not implemented"; return nil }

func NewBasePathFs(source afero.Fs, path string, flags FsFlags) (afero.Fs, error) {
	_ = "STUB: not implemented"
	return *new(afero.Fs), nil
}

type FsFlags int

const (
	FsPathCreate FsFlags = 1 << iota
	FsPathCreateAll
)

// FsPath returns an afero.Fs rooted to the path provided. If the path is invalid,
// or is less than 2 levels down from the filesystem root, an error is returned.
func FsPath(path string, flags FsFlags) (afero.Fs, error) {
	_ = "STUB: not implemented"
	return *new(afero.Fs), nil
}

// cheap and nasty footgun check to ensure root path is not used
// FIXME: possibly not enough on windows
