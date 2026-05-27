package s3bolt

import (
	"io"

	bolt "go.etcd.io/bbolt"

	"github.com/johannesboyne/gofakes3"
)

var (
	emptyPrefix = &gofakes3.Prefix{}
)

type Backend struct {
	bolt           *bolt.DB
	timeSource     gofakes3.TimeSource
	metaBucketName []byte
}

var _ gofakes3.Backend = &Backend{}

type Option func(b *Backend)

func WithTimeSource(timeSource gofakes3.TimeSource) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewFile(file string, opts ...Option) (*Backend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(bolt *bolt.DB, opts ...Option) *Backend { _ = "STUB: not implemented"; return nil }

// Underscore guarantees no overlap with legal S3 bucket names

// metaBucket returns a utility that manages access to the metadata bucket.
// The returned struct is valid only for the lifetime of the bolt.Tx.
// The metadata bucket may not exist if this is an older database.
func (db *Backend) metaBucket(tx *bolt.Tx) (*metaBucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: support legacy databases; remove when versioning is supported.

func (db *Backend) ListBuckets() ([]gofakes3.BucketInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempt to assign metadata. If it isn't found, we will just
// pretend that's fine for now. This is to support existing
// databases that have buckets created without associated metadata.
//
// FIXME: clean this up when there is an upgrade script to expect
// that it exists

// The AWS CLI will fail if there is no creation date:

func (db *Backend) ListBucket(name string, prefix *gofakes3.Prefix, page gofakes3.ListBucketPage) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Backend) CreateBucket(name string) error { _ = "STUB: not implemented"; return nil }

// create bucket metadata

// create bucket

func (db *Backend) DeleteBucket(name string) error { _ = "STUB: not implemented"; return nil }

// delete bucket

// delete bucket metadata

// FIXME: assumes a legacy database, where the bucket may not exist. Clean
// this up when there is a DB upgrade script.

func (db *Backend) ForceDeleteBucket(name string) error { _ = "STUB: not implemented"; return nil }

// Delete all objects in the bucket

// Delete bucket metadata

// FIXME: assumes a legacy database, where the bucket may not exist. Clean
// this up when there is a DB upgrade script.

// Delete the bucket itself

func (db *Backend) BucketExists(name string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (db *Backend) HeadObject(bucketName, objectName string) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Backend) GetObject(bucketName, objectName string, rangeRequest *gofakes3.ObjectRangeRequest) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: objectName here is a bit of a hack; this can be cleaned up when we have a
// database migration script.

func (db *Backend) PutObject(
	bucketName, objectName string,
	meta map[string]string,
	input io.Reader,
	size int64,
	conditions *gofakes3.PutConditions,
) (result gofakes3.PutObjectResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.PutObjectResult), nil
}

func (db *Backend) CopyObject(srcBucket, srcKey, dstBucket, dstKey string, meta map[string]string) (result gofakes3.CopyObjectResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.CopyObjectResult), nil
}

func (db *Backend) DeleteObject(bucketName, objectName string) (result gofakes3.ObjectDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.ObjectDeleteResult), nil
}

func (db *Backend) DeleteMulti(bucketName string, objects ...string) (result gofakes3.MultiDeleteResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.MultiDeleteResult), nil
}

// getConditionalObjectInfo returns information about an object for conditional checking.
// This method assumes it's called within a bolt transaction.
func (db *Backend) getConditionalObjectInfo(bucket *bolt.Bucket, objectName string) (*gofakes3.ConditionalObjectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
