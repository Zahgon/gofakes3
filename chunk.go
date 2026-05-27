package gofakes3

import (
	"io"
)

type chunkedReader struct {
	inner         io.Reader
	chunkRemain   int
	notFirstChunk bool
}

func newChunkedReader(inner io.Reader) *chunkedReader { _ = "STUB: not implemented"; return nil }

func (r *chunkedReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// read sizeToRead bytes from inner reader
// to p, start from n.
// n is bytes already read.

// read until this chunk ends

// Is first chunk.

// skip last chunk's b"\r\n"

// read next chunk header

// "chunk-signature=" + sizeOfHash + "\r\n"
