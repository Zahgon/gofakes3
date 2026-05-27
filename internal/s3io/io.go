package s3io

import "io"

type ReaderWithDummyCloser struct{ io.Reader }

func (d ReaderWithDummyCloser) Close() error { _ = "STUB: not implemented"; return nil }

type NoOpReadCloser struct{}

func (d NoOpReadCloser) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d NoOpReadCloser) Close() error { _ = "STUB: not implemented"; return nil }
