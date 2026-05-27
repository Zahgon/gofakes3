package gofakes3

import (
	"hash"
	"io"
)

// hashingReader proxies an existing io.Reader, passing each read block to the
// given hash.Hash.
//
// If the expected hash is not empty, once the underlying reader returns EOF,
// the hash is checked.
type hashingReader struct {
	inner    io.Reader
	expected []byte
	hash     hash.Hash
	sum      []byte
}

func newHashingReader(inner io.Reader, expectedMD5Base64 string) (*hashingReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sum returns the hash of the data read from the inner reader so far.
// If into is passed, it may be used if the hash needs to be computed.
func (h *hashingReader) Sum(into []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *hashingReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Hash.Write never returns an error.

// FIXME: some more context here would be useful; need to flush out
// what S3 responds with in this case.
