package s3afero

import (
	"io"
	"sync"
	"time"

	"github.com/spf13/afero"

	"github.com/johannesboyne/gofakes3"
)

// SingleBucketBackend is a gofakes3.Backend that allows you to treat an existing
// filesystem as an S3 bucket directly. It does not support multiple buckets.
//
// A second afero.Fs, metaFs, may be passed; if this is nil,
// afero.NewMemMapFs() is used and the metadata will not persist between
// restarts of gofakes3.
//
// It is STRONGLY recommended that the metadata Fs is not contained within the
// `/buckets` subdirectory as that could make a significant mess, but this is
// infeasible to validate, so you're encouraged to be extremely careful!
type SingleBucketBackend struct {
	lock      sync.Mutex
	fs        afero.Fs
	metaStore *metaStore
	name      string
}

var _ gofakes3.Backend = &SingleBucketBackend{}

func SingleBucket(name string, fs afero.Fs, metaFs afero.Fs, opts ...SingleOption) (*SingleBucketBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *SingleBucketBackend) ListBuckets() ([]gofakes3.BucketInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: "birth time" is not available cross-platform.
// See MultiBucketBackend.ListBuckets for more details.

func (db *SingleBucketBackend) ListBucket(bucket string, prefix *gofakes3.Prefix, page gofakes3.ListBucketPage) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *SingleBucketBackend) getBucketWithFilePrefixLocked(bucket string, prefixPath, prefixPart string) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expected use of 'path'; see the "Path Handling" subheading in doc.go:

func (db *SingleBucketBackend) getBucketWithArbitraryPrefixLocked(bucket string, prefix *gofakes3.Prefix) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *SingleBucketBackend) ensureMeta(
	bucket string,
	objectPath string,
	size int64,
	mtime time.Time,
) (meta *Metadata, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *SingleBucketBackend) HeadObject(bucketName, objectName string) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *SingleBucketBackend) GetObject(bucketName, objectName string, rangeRequest *gofakes3.ObjectRangeRequest) (obj *gofakes3.Object, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If an error occurs, the caller may not have access to Object.Body in order to close it:

func (db *SingleBucketBackend) PutObject(
	bucketName, objectName string,
	meta map[string]string,
	input io.Reader,
	size int64,
	conditions *gofakes3.PutConditions,
) (result gofakes3.PutObjectResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.PutObjectResult), nil
}

// Unfortunately, afero's MemMapFs updates the mtime if you double-close, which
// highlights that other afero.Fs implementations may have side effects here::

// We have to close here before we stat the file as some filesystems don't update the
// mtime until after close:

func (db *SingleBucketBackend) DeleteMulti(bucketName string, objects ...string) (result gofakes3.MultiDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.MultiDeleteResult), nil
}

func (db *SingleBucketBackend) CopyObject(srcBucket, srcKey, dstBucket, dstKey string, meta map[string]string) (result gofakes3.CopyObjectResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.CopyObjectResult), nil
}

func (db *SingleBucketBackend) DeleteObject(bucketName, objectName string) (result gofakes3.ObjectDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.ObjectDeleteResult), nil
}

func (db *SingleBucketBackend) deleteObjectLocked(bucketName, objectName string) error {
	_ = "STUB: not implemented"
	// S3 does not report an error when attemping to delete a key that does not exist, so
	// we need to skip IsNotExist errors.
	return nil
}

// CreateBucket cannot be implemented by this backend. See MultiBucketBackend if you
// need a backend that supports it.
func (db *SingleBucketBackend) CreateBucket(name string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteBucket cannot be implemented by this backend. See MultiBucketBackend if you
// need a backend that supports it.
func (db *SingleBucketBackend) DeleteBucket(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *SingleBucketBackend) ForceDeleteBucket(name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete all objects in the bucket

// Delete the bucket itself

func (db *SingleBucketBackend) BucketExists(name string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false,

		// getConditionalObjectInfo returns information about an object for conditional checking.
		// This method assumes the backend lock is already held.
		nil
}

func (db *SingleBucketBackend) getConditionalObjectInfo(bucketName, objectName string) (*gofakes3.ConditionalObjectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
