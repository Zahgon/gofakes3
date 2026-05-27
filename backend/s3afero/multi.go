package s3afero

import (
	"io"
	"os"
	"sync"

	"github.com/spf13/afero"

	"github.com/johannesboyne/gofakes3"
)

// MultiBucketBackend is a gofakes3.Backend that allows you to create multiple
// buckets within the same afero.Fs. Buckets are stored under the `/buckets`
// subdirectory. Metadata is stored in the `/metadata` subdirectory by default,
// but any afero.Fs can be used.
//
// It is STRONGLY recommended that the metadata Fs is not contained within the
// `/buckets` subdirectory as that could make a significant mess, but this is
// infeasible to validate, so you're encouraged to be extremely careful!
type MultiBucketBackend struct {
	lock      sync.Mutex
	baseFs    afero.Fs
	bucketFs  afero.Fs
	metaStore *metaStore
	dirMode   os.FileMode
	flags     FsFlags

	// FIXME(bw): values in here should not be used beyond the configuration
	// step; maybe this can be cleaned up later using a builder struct or
	// something.
	configOnly struct {
		metaFs afero.Fs
	}
}

var _ gofakes3.Backend = &MultiBucketBackend{}

func MultiBucket(fs afero.Fs, opts ...MultiOption) (*MultiBucketBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *MultiBucketBackend) ListBuckets() ([]gofakes3.BucketInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: "birth time" is not available cross-platform.
// https://github.com/djherbis/times provides access to it on supported
// platforms, but that wouldn't really be compatible with afero.
// ModTime and some documented caveats might be the least-worst
// option for this particular backend:

func (db *MultiBucketBackend) ListBucket(bucket string, prefix *gofakes3.Prefix, page gofakes3.ListBucketPage) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *MultiBucketBackend) getBucketWithFilePrefixLocked(bucket string, prefixPath, prefixPart string) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expected use of 'path'; see the "Path Handling" subheading in doc.go:

func (db *MultiBucketBackend) getBucketWithArbitraryPrefixLocked(bucket string, prefix *gofakes3.Prefix) (*gofakes3.ObjectList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// should never happen

func (db *MultiBucketBackend) CreateBucket(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *MultiBucketBackend) DeleteBucket(name string) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// This check is slightly racy. If another service outside gofakes3
// changes the filesystem between this check and the call to Remove,
// the bucket may be deleted even though there are items in it. You
// would expect that afero.Fs would raise an error if you tried to
// delete a directory that had stuff in it, but implementers of
// afero.Fs may not implement that particular constraint. We have no
// choice but to fall back on the db's lock and assume that a race
// won't happen.

// FIXME(bw): the error handling logic here is a little janky:

func (db *MultiBucketBackend) ForceDeleteBucket(name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete all objects in the bucket

// Delete the bucket itself

// Delete bucket metadata

func (db *MultiBucketBackend) BucketExists(name string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (db *MultiBucketBackend) HeadObject(bucketName, objectName string) (*gofakes3.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Another slighly racy check:

func (db *MultiBucketBackend) GetObject(bucketName, objectName string, rangeRequest *gofakes3.ObjectRangeRequest) (obj *gofakes3.Object, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Another slighly racy check:

// If an error occurs, the caller may not have access to Object.Body in order to close it:

func (db *MultiBucketBackend) PutObject(
	bucketName, objectName string,
	meta map[string]string,
	input io.Reader,
	size int64,
	conditions *gofakes3.PutConditions,
) (result gofakes3.PutObjectResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.PutObjectResult), nil
}

// Another slighly racy check:

// Unfortunately, afero's MemMapFs updates the mtime if you double-close, which
// highlights that other afero.Fs implementations may have side effects here::

// We have to close here before we stat the file as some filesystems don't update the
// mtime until after close:

func (db *MultiBucketBackend) CopyObject(srcBucket, srcKey, dstBucket, dstKey string, meta map[string]string) (result gofakes3.CopyObjectResult, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.CopyObjectResult), nil
}

func (db *MultiBucketBackend) DeleteObject(bucketName, objectName string) (result gofakes3.ObjectDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.ObjectDeleteResult), nil
}

// Another slighly racy check:

func (db *MultiBucketBackend) deleteObjectLocked(bucketName, objectName string) error {
	_ = "STUB: not implemented"
	return nil
}

// S3 does not report an error when attemping to delete a key that does not exist, so
// we need to skip IsNotExist errors.

func (db *MultiBucketBackend) DeleteMulti(bucketName string, objects ...string) (result gofakes3.MultiDeleteResult, rerr error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.MultiDeleteResult), nil
}

// Another slighly racy check:

// getConditionalObjectInfo returns information about an object for conditional checking.
// This method assumes the backend lock is already held.
func (db *MultiBucketBackend) getConditionalObjectInfo(bucketName, objectName string) (*gofakes3.ConditionalObjectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
