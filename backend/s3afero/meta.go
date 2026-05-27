package s3afero

import (
	"time"

	"github.com/spf13/afero"
)

type Metadata struct {
	File    string
	ModTime time.Time
	Size    int64
	Hash    []byte
	Meta    map[string]string
}

type metaPath struct {
	bucket string
	object string
}

func (mp metaPath) FilePath() string { _ = "STUB: not implemented"; return "" }

type metaStore struct {
	fs          afero.Fs
	modTimeCalc modTimeCalc
	modTimeRes  time.Duration
}

func newMetaStore(fs afero.Fs, modTimeCalc modTimeCalc) *metaStore {
	_ = "STUB: not implemented"
	return nil
}

func (ms *metaStore) getModTimeRes() (dur time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (ms *metaStore) metaPath(bucket string, object string) metaPath {
	_ = "STUB: not implemented"
	// FIXME: may need to add path segments but that may be a thing of the past:
	// https://stackoverflow.com/questions/466521/how-many-files-can-i-put-in-a-directory
	return *new(metaPath)
}

func (ms *metaStore) loadMeta(bucket string, object string, size int64, mtime time.Time) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ms *metaStore) saveMeta(path metaPath, meta *Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *metaStore) deleteMeta(path metaPath) error { _ = "STUB: not implemented"; return nil }

func (ms *metaStore) deleteBucket(bucket string) error { _ = "STUB: not implemented"; return nil }
