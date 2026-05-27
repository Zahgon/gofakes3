package s3mem

import (
	"time"

	"github.com/johannesboyne/gofakes3"
	"github.com/ryszard/goskiplist/skiplist"
)

type versionGenFunc func() gofakes3.VersionID

type versioningStatus int

type bucket struct {
	name         string
	versioning   gofakes3.VersioningStatus
	versionGen   versionGenFunc
	creationDate gofakes3.ContentTime

	objects *skiplist.SkipList
}

func newBucket(name string, at time.Time, versionGen versionGenFunc) *bucket {
	_ = "STUB: not implemented"
	return nil
}

type bucketObject struct {
	name     string
	data     *bucketData
	versions *skiplist.SkipList
}

func (b *bucketObject) Iterator() *bucketObjectIterator { _ = "STUB: not implemented"; return nil }

type bucketObjectIterator struct {
	data     *bucketData
	iter     skiplist.Iterator
	cur      *bucketData
	seenData bool
	done     bool
}

func (b *bucketObjectIterator) Seek(key gofakes3.VersionID) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *bucketObjectIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (b *bucketObjectIterator) Close() { _ = "STUB: not implemented"; return }

func (b *bucketObjectIterator) Value() *bucketData { _ = "STUB: not implemented"; return nil }

type bucketData struct {
	name         string
	lastModified time.Time
	versionID    gofakes3.VersionID
	deleteMarker bool
	body         []byte
	hash         []byte
	metadata     map[string]string
}

func (bi *bucketData) toObject(rangeRequest *gofakes3.ObjectRangeRequest, withBody bool) (obj *gofakes3.Object, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In case of a range request the correct part of the slice is extracted:

// The data slice should be completely replaced if the bucket item is edited, so
// it should be safe to return the data slice directly.

func (b *bucket) setVersioning(enabled bool) { _ = "STUB: not implemented"; return }

func (b *bucket) object(objectName string) (obj *bucketObject) {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) objectVersion(objectName string, versionID gofakes3.VersionID) (*bucketData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) put(name string, item *bucketData) {
	_ = "STUB: not implemented"
	// Always generate a version for convenience; we can just mask it on return.
	return
}

func (b *bucket) rm(name string, at time.Time) (result gofakes3.ObjectDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.ObjectDeleteResult), nil
}

// S3 does not report an error when attemping to delete a key that does not exist

func (b *bucket) rmVersion(name string, versionID gofakes3.VersionID, at time.Time) (result gofakes3.ObjectDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.ObjectDeleteResult), nil
}

// S3 does not report an error when attemping to delete a key that does not exist
