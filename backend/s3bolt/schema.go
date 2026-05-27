package s3bolt

// The schema for the bolt database is described in here. External users of the
// database should consider this an internal implementation detail, subject to
// change without notice or version number changes.
//
// This may change in the future.

import (
	"time"

	"github.com/johannesboyne/gofakes3"
	bolt "go.etcd.io/bbolt"
)

type boltBucket struct {
	CreationDate time.Time
}

type boltObject struct {
	Name         string
	Metadata     map[string]string
	LastModified time.Time
	Size         int64
	Contents     []byte
	Hash         []byte
}

func (b *boltObject) Object(objectName string, rangeRequest *gofakes3.ObjectRangeRequest) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bucketMetaKey(name string) []byte { _ = "STUB: not implemented"; return nil }

type metaBucket struct {
	*bolt.Tx
	metaName []byte
	bucket   *bolt.Bucket
}

func (mb *metaBucket) deleteS3Bucket(bucket string) error { _ = "STUB: not implemented"; return nil }

func (mb *metaBucket) createS3Bucket(bucket string, at time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (mb *metaBucket) s3Bucket(bucket string) (*boltBucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: should return an error once database upgrades are supported.
