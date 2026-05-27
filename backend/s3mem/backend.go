package s3mem

import (
	"io"
	"sync"

	"github.com/johannesboyne/gofakes3"
)

var (
	emptyPrefix       = &gofakes3.Prefix{}
	emptyVersionsPage = &gofakes3.ListBucketVersionsPage{}
)

type Backend struct {
	buckets          map[string]*bucket
	timeSource       gofakes3.TimeSource
	versionGenerator *versionGenerator
	versionSeed      int64
	versionSeedSet   bool
	versionScratch   []byte
	lock             sync.RWMutex
}

var _ gofakes3.Backend = &Backend{}
var _ gofakes3.VersionedBackend = &Backend{}

type Option func(b *Backend)

func WithTimeSource(timeSource gofakes3.TimeSource) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithVersionSeed(seed int64) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(opts ...Option) *Backend { _ = "STUB: not implemented"; return nil }

func (db *Backend) ListBuckets() ([]gofakes3.BucketInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Backend) ListBucket(name string, prefix *gofakes3.Prefix, page gofakes3.ListBucketPage) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the current item is the Marker, move to the next item.

// Should not count towards keys

func (db *Backend) CreateBucket(name string) error { _ = "STUB: not implemented"; return nil }

func (db *Backend) DeleteBucket(name string) error { _ = "STUB: not implemented"; return nil }

func (db *Backend) ForceDeleteBucket(name string) error { _ = "STUB: not implemented"; return nil }

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

// FIXME: If the current version of the object is a delete marker,
// Amazon S3 behaves as if the object was deleted and includes
// x-amz-delete-marker: true in the response.
//
// The solution may be to return an object but no error if the object is
// a delete marker, and let the main GoFakeS3 class decide what to do.

func (db *Backend) PutObject(
	bucketName,
	objectName string,
	meta map[string]string,
	input io.Reader,
	size int64,
	conditions *gofakes3.PutConditions,
) (result gofakes3.PutObjectResult, err error) {
	_ = "STUB: not implemented"
	// No need to lock the backend while we read the data into memory; it holds
	// the write lock open unnecessarily, and could be blocked for an unreasonably
	// long time by a connection timing out:
	return *new(gofakes3.PutObjectResult), nil
}

// versionID is assigned in bucket.put()

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

// FIXME: what to do with rm result in multi delete?

// FIXME: log

func (db *Backend) DeleteMultiVersions(bucketName string, objects ...gofakes3.ObjectID) (result gofakes3.MultiDeleteResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.MultiDeleteResult), nil
}

// FIXME: what to do with rm result in multi delete?

// FIXME: log

func (db *Backend) VersioningConfiguration(bucketName string) (versioning gofakes3.VersioningConfiguration, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.VersioningConfiguration), nil
}

func (db *Backend) SetVersioningConfiguration(bucketName string, v gofakes3.VersioningConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *Backend) GetObjectVersion(
	bucketName, objectName string,
	versionID gofakes3.VersionID,
	rangeRequest *gofakes3.ObjectRangeRequest) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Backend) HeadObjectVersion(bucketName, objectName string, versionID gofakes3.VersionID) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Backend) DeleteObjectVersion(bucketName, objectName string, versionID gofakes3.VersionID) (result gofakes3.ObjectDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.ObjectDeleteResult), nil
}

func (db *Backend) ListBucketVersions(
	bucketName string,
	prefix *gofakes3.Prefix,
	page *gofakes3.ListBucketVersionsPage,
) (*gofakes3.ListBucketVersionsResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: NO idea what S3 would do here.

// FIXME: The S3 docs have this to say on the topic of result ordering:
//   "The following request returns objects in the order they were stored,
//   returning the most recently stored object first starting with the value
//   for key-marker."
//
// OK so this method....
// - Returns objects in the order they were stored
// - Returning the most recently stored object first
//
// This makes no sense at all!

// FIXME: log

// S300005

// S300005

// getConditionalObjectInfo returns information about an object for conditional checking.
// This method assumes the bucket lock is already held.
func (db *Backend) getConditionalObjectInfo(bucket *bucket, objectName string) (*gofakes3.ConditionalObjectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nextVersion assumes the backend's lock is acquired
func (db *Backend) nextVersion() gofakes3.VersionID {
	_ = "STUB: not implemented"
	return *new(gofakes3.VersionID)
}
