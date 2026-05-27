package s3mem

import (
	"math/big"
	"sync"

	"github.com/johannesboyne/gofakes3"
)

var add1 = new(big.Int).SetInt64(1)

type versionGenerator struct {
	state uint64
	size  int
	next  *big.Int
	mu    sync.Mutex
}

func newVersionGenerator(seed uint64, size int) *versionGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (v *versionGenerator) Next(scratch []byte) (gofakes3.VersionID, []byte) {
	_ = "STUB: not implemented"
	return *new(gofakes3.VersionID), nil
}

// cheap and nasty way to ensure a multiple of 8 definitely greater than size

// This is a simple inline implementation of http://xoshiro.di.unimi.it/splitmix64.c.
// It may not ultimately be the right tool for this job but with a large
// enough size the collision risk should still be minuscule.

// The version IDs that come out of S3 appear to start with '3/' and follow
// with a base64-URL encoded blast of god knows what. There didn't appear
// to be any explanation of the format beyond that, but let's copy it anyway.
//
// Base64 is not sortable though, and we need our versions to be lexicographically
// sortable for the SkipList key, so we have to encode it as base32hex, which _is_
// sortable, and just pretend that it's "Base64". Phew!
